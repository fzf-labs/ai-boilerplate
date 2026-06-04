package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetPaymentInfo 获取支付信息接口-选择支付方式并获取支付信息
func (a *AppV1MallOrderService) GetPaymentInfo(ctx context.Context, req *pb.GetPaymentInfoReq) (*pb.GetPaymentInfoReply, error) {
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
		return nil, pb.ErrorReasonParamError() // 订单已支付或已取消
	}

	// 5. 检查订单是否过期
	if order.ExpiredTime.Valid && time.Now().After(order.ExpiredTime.Time) {
		return nil, pb.ErrorReasonParamError() // 订单已过期
	}

	_ = req
	return nil, pb.ErrorReasonAPIThirdErr(pb.WithFmtMsg("支付功能未接入"))
}
