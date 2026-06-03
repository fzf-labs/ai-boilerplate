package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"gorm.io/datatypes"
)

// PaymentCallback 支付回调接口
func (a *AppV1MallOrderService) PaymentCallback(ctx context.Context, req *pb.PaymentCallbackReq) (*pb.PaymentCallbackReply, error) {
	resp := &pb.PaymentCallbackReply{
		Success: false,
		Message: "处理失败",
	}

	// 1. 查询订单
	order, err := a.mallOrderRepo.FindOneCacheByID(ctx, req.GetOrderId())
	if err != nil {
		resp.Message = "订单查询失败"
		return resp, nil
	}
	if order == nil || order.ID == "" {
		resp.Message = "订单不存在"
		return resp, nil
	}
	if req.GetPaymentStatus() != 1 && req.GetPaymentStatus() != 2 {
		resp.Message = "支付状态无效"
		return resp, nil
	}

	// 2. 校验订单状态
	if order.PaymentStatus == 1 {
		resp.Success = true
		resp.Message = "订单已支付"
		return resp, nil
	}

	now := time.Now()
	transactionID := strings.TrimSpace(req.GetTransactionId())
	if transactionID == "" {
		transactionID = order.ID
	}
	paymentMethod := strings.TrimSpace(req.GetPaymentMethod())
	if paymentMethod == "" {
		paymentMethod = order.PaymentMethod
	}
	callbackJSON := datatypes.JSON([]byte("{}"))
	rawCallbackData := strings.TrimSpace(req.GetCallbackData())
	if rawCallbackData != "" {
		if json.Valid([]byte(rawCallbackData)) {
			callbackJSON = datatypes.JSON([]byte(rawCallbackData))
		} else {
			encoded, _ := json.Marshal(map[string]string{"raw": rawCallbackData})
			callbackJSON = datatypes.JSON(encoded)
		}
	}
	paymentRecord, err := a.mallPaymentRecordRepo.FindOneCacheByTransactionID(ctx, transactionID)
	if err != nil {
		resp.Message = "支付记录查询失败"
		return resp, nil
	}

	err = a.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		oldOrder := a.mallOrderRepo.DeepCopy(order)
		order.PaymentMethod = paymentMethod
		if req.GetPaymentStatus() == 1 {
			order.PaymentStatus = 1
			order.PaymentTime = sql.NullTime{Time: now, Valid: true}
			if order.ProductType == constant.MallProductTypeMembership.String() {
				order.Status = "completed"
				order.DeliveryTime = sql.NullTime{Time: now, Valid: true}
			} else {
				order.Status = "pendingDelivery"
			}
		} else {
			order.PaymentStatus = 2
			order.Status = "canceled"
		}
		if err := a.mallOrderRepo.UpdateOneCacheWithZeroByTx(ctx, tx, order, oldOrder); err != nil {
			return err
		}

		if paymentRecord == nil || paymentRecord.ID == "" {
			paymentRecord = a.mallPaymentRecordRepo.NewData()
			paymentRecord.OrderID = order.ID
			paymentRecord.TransactionID = transactionID
			paymentRecord.PaymentChannel = paymentMethod
			paymentRecord.PaymentMethod = paymentMethod
			paymentRecord.Amount = order.ActualAmount
			paymentRecord.Currency = order.Currency
			paymentRecord.Status = int32(constant.StatusEnable)
			paymentRecord.ThirdPartyOrderNo = order.ID
			paymentRecord.CreatedAt = now
			paymentRecord.UpdatedAt = now
			paymentRecord.PaymentStatus = req.GetPaymentStatus()
			paymentRecord.ThirdPartyTransactionID = strings.TrimSpace(req.GetTransactionId())
			paymentRecord.CallbackData = callbackJSON
			paymentRecord.CallbackTime = sql.NullTime{Time: now, Valid: true}
			if req.GetPaymentStatus() == 2 {
				paymentRecord.ErrorMessage = "支付失败"
			}
			if err := a.mallPaymentRecordRepo.CreateOneCacheByTx(ctx, tx, paymentRecord); err != nil {
				return err
			}
		} else {
			oldPaymentRecord := a.mallPaymentRecordRepo.DeepCopy(paymentRecord)
			paymentRecord.PaymentChannel = paymentMethod
			paymentRecord.PaymentMethod = paymentMethod
			paymentRecord.Amount = order.ActualAmount
			paymentRecord.Currency = order.Currency
			paymentRecord.PaymentStatus = req.GetPaymentStatus()
			paymentRecord.ThirdPartyTransactionID = strings.TrimSpace(req.GetTransactionId())
			paymentRecord.CallbackData = callbackJSON
			paymentRecord.CallbackTime = sql.NullTime{Time: now, Valid: true}
			paymentRecord.Status = int32(constant.StatusEnable)
			if req.GetPaymentStatus() == 1 {
				paymentRecord.ErrorMessage = ""
			} else {
				paymentRecord.ErrorMessage = "支付失败"
			}
			if err := a.mallPaymentRecordRepo.UpdateOneCacheWithZeroByTx(ctx, tx, paymentRecord, oldPaymentRecord); err != nil {
				return err
			}
		}

		if req.GetPaymentStatus() != 1 || order.ProductType != constant.MallProductTypeMembership.String() {
			return nil
		}

		product, err := a.mallProductRepo.FindOneCacheByID(ctx, order.ProductID)
		if err != nil {
			return err
		}
		if product == nil || product.ID == "" {
			return pb.ErrorReasonDataRecordNotFound()
		}
		productConfig, err := a.mallProductRepo.GetMembershipProductConfig(product)
		if err != nil {
			return err
		}
		if productConfig.Membership.MembershipType == "" || productConfig.Membership.DurationDays <= 0 {
			return pb.ErrorReasonActivationCodeProductConfigInvalid()
		}
		userMembership, err := a.userMembershipRepo.GetUserActualMembershipInfo(ctx, order.UserID)
		if err != nil {
			return err
		}
		oldUserMembership, _ := json.Marshal(userMembership)
		oldMembershipData := a.userMembershipRepo.DeepCopy(userMembership)
		newMembershipType, newExpiredAt, err := a.userMembershipRepo.CalcMembershipChange(
			ctx,
			userMembership.MembershipType,
			userMembership.ExpiredAt.Time,
			productConfig.Membership.MembershipType,
			int(productConfig.Membership.DurationDays),
		)
		if err != nil {
			return err
		}
		userMembership.MembershipType = newMembershipType
		userMembership.ExpiredAt = sql.NullTime{Time: newExpiredAt, Valid: !newExpiredAt.IsZero()}
		userMembership.Status = int32(constant.StatusEnable)
		if err := a.userMembershipRepo.UpdateOneCacheWithZeroByTx(ctx, tx, userMembership, oldMembershipData); err != nil {
			return err
		}
		newUserMembership, _ := json.Marshal(userMembership)
		changeRecord := &ai_boilerplate_model.UserMembershipChange{
			UserID:     order.UserID,
			SourceType: constant.MembershipChangeSourceOrder.String(),
			SourceID:   order.ID,
			Before:     oldUserMembership,
			After:      newUserMembership,
			Remark:     "会员订单支付成功",
		}
		if err := a.userMembershipChangeRepo.CreateOneCacheByTx(ctx, tx, changeRecord); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		resp.Success = false
		resp.Message = "订单更新失败"
		return resp, nil
	}
	resp.Success = true
	if req.GetPaymentStatus() == 1 {
		resp.Message = "支付成功"
	} else {
		resp.Message = "支付失败,订单已取消"
	}
	return resp, nil
}
