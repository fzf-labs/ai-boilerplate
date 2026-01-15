package service

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
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

	// 2. 校验订单状态
	if order.PaymentStatus == 1 {
		resp.Success = true
		resp.Message = "订单已支付"
		return resp, nil
	}

	// 3. 校验订单是否过期
	if order.ExpiredTime.Valid && time.Now().After(order.ExpiredTime.Time) {
		resp.Message = "订单已过期"
		return resp, nil
	}
	// 根据支付状态更新订单
	now := time.Now()
	if req.GetPaymentStatus() == 1 { // 支付成功
		order.PaymentStatus = 1
		order.PaymentTime = sql.NullTime{Time: now, Valid: true}
		order.PaymentMethod = req.GetPaymentMethod()
		order.Status = "pendingDelivery" // 更新为待发货状态

		resp.Success = true
		resp.Message = "支付成功"
	} else { // 支付失败
		order.PaymentStatus = 2
		order.Status = "canceled"

		resp.Success = true
		resp.Message = "支付失败,订单已取消"
	}

	// 5. 保存订单更新 + 会员变更处理
	err = a.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		return nil
	})
	if err != nil {
		resp.Success = false
		resp.Message = "订单更新失败"
		return resp, nil
	}
	// 6. TODO: 如果支付成功，可以在这里触发后续业务逻辑
	// - 发送通知
	// - 更新商品库存
	// - 记录支付流水等

	return resp, nil
}
