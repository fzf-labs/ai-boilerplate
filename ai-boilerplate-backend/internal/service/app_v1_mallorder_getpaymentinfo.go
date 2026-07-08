package service

import (
	"context"
	"strconv"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetPaymentInfo 获取支付信息接口-选择支付方式并获取支付信息
func (a *AppV1MallOrderService) GetPaymentInfo(ctx context.Context, req *pb.GetPaymentInfoReq) (*pb.GetPaymentInfoReply, error) {
	resp := &pb.GetPaymentInfoReply{}

	// 1. 获取当前用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 2. 查询订单
	order, err := a.mallOrderRepo.FindOneCacheByID(ctx, req.GetOrderId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if order == nil || order.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	// 3. 校验订单归属
	if order.UserID != userID {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 4. 检查订单状态（必须是待支付状态）
	if order.PaymentStatus != 0 {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("订单已支付或已关闭"))
	}

	now := time.Now()
	// 5. 检查订单是否过期
	if order.ExpiredTime.Valid && now.After(order.ExpiredTime.Time) {
		oldOrder := a.mallOrderRepo.DeepCopy(order)
		order.PaymentStatus = 2
		order.Status = "canceled"
		if err := a.mallOrderRepo.UpdateOneCacheWithZero(ctx, order, oldOrder); err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("订单已过期"))
	}

	// 6. 首个支付渠道明确为微信支付 mock/sandbox 通道
	if !isSupportedPaymentChannel(req.GetPaymentMethod()) {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("暂仅支持微信支付"))
	}

	transactionID := buildMockPaymentTransactionID(order.ID)
	paymentRecord, err := a.mallPaymentRecordRepo.FindOneCacheByTransactionID(ctx, transactionID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if paymentRecord != nil && paymentRecord.ID != "" && paymentRecord.PaymentStatus == 1 {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("订单已支付"))
	}

	err = a.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		if order.PaymentMethod != paymentChannelWechat {
			oldOrder := a.mallOrderRepo.DeepCopy(order)
			order.PaymentMethod = paymentChannelWechat
			if err := a.mallOrderRepo.UpdateOneCacheWithZeroByTx(ctx, tx, order, oldOrder); err != nil {
				return err
			}
		}

		if paymentRecord == nil || paymentRecord.ID == "" {
			paymentRecord = a.mallPaymentRecordRepo.NewData()
			paymentRecord.OrderID = order.ID
			paymentRecord.TransactionID = transactionID
			paymentRecord.ThirdPartyOrderNo = order.ID
			paymentRecord.CreatedAt = now
		} else {
			paymentRecord = a.mallPaymentRecordRepo.DeepCopy(paymentRecord)
		}
		paymentRecord.PaymentChannel = paymentChannelWechat
		paymentRecord.PaymentMethod = paymentMethodH5
		paymentRecord.Amount = order.ActualAmount
		paymentRecord.Currency = order.Currency
		paymentRecord.PaymentStatus = 0
		paymentRecord.Status = int32(constant.StatusEnable)
		paymentRecord.ErrorCode = ""
		paymentRecord.ErrorMessage = ""
		paymentRecord.UpdatedAt = now

		if paymentRecord.ID == "" {
			return a.mallPaymentRecordRepo.CreateOneCacheByTx(ctx, tx, paymentRecord)
		}
		oldPaymentRecord, err := a.mallPaymentRecordRepo.FindOneCacheByTransactionID(ctx, transactionID)
		if err != nil {
			return err
		}
		return a.mallPaymentRecordRepo.UpdateOneCacheWithZeroByTx(ctx, tx, paymentRecord, oldPaymentRecord)
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	timestamp := strconv.FormatInt(now.Unix(), 10)
	nonce := strconv.FormatInt(now.UnixNano(), 36)
	paySign := buildMockPaymentSign(order.ID, transactionID, 0, order.ActualAmount, timestamp, nonce)

	resp.OrderId = order.ID
	resp.PaymentMethod = paymentChannelWechat
	resp.ActualAmount = order.ActualAmount
	resp.PaymentUrl = buildMockPaymentURL(order, transactionID, timestamp, nonce, paySign)
	resp.QrCodeUrl = resp.PaymentUrl
	resp.PrepayId = transactionID
	if order.ExpiredTime.Valid {
		resp.ExpiredTime = order.ExpiredTime.Time.Format(time.RFC3339)
	}
	resp.AppId = mockPaymentAppID
	resp.TimeStamp = timestamp
	resp.NonceStr = nonce
	resp.SignType = mockPaymentSignType
	resp.PaySign = paySign

	return resp, nil
}
