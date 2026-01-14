package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetOrderInfo 查询订单接口-单条订单查询
func (a *AppV1MallOrderService) GetOrderInfo(ctx context.Context, req *pb.GetOrderInfoReq) (*pb.GetOrderInfoReply, error) {
	resp := &pb.GetOrderInfoReply{}

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

	// 3. 校验订单归属（确保用户只能查询自己的订单）
	if order.UserID != userID {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 4. 构建返回信息
	resp.Info = &pb.MallOrderInfo{
		Id:             order.ID,
		UserId:         order.UserID,
		ProductType:    order.ProductType,
		ProductId:      order.ProductID,
		OriginalAmount: order.OriginalAmount,
		DiscountAmount: order.DiscountAmount,
		ActualAmount:   order.ActualAmount,
		RefundAmount:   order.RefundAmount,
		Currency:       order.Currency,
		PaymentMethod:  order.PaymentMethod,
		PaymentStatus:  order.PaymentStatus,
		Status:         order.Status,
		Remark:         order.Remark,
		CreatedAt:      order.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      order.UpdatedAt.Format(time.RFC3339),
	}

	// 填充时间字段
	if order.PaymentTime.Valid {
		resp.Info.PaymentTime = order.PaymentTime.Time.Format(time.RFC3339)
	}
	if order.DeliveryTime.Valid {
		resp.Info.DeliveryTime = order.DeliveryTime.Time.Format(time.RFC3339)
	}
	if order.ExpiredTime.Valid {
		resp.Info.ExpiredTime = order.ExpiredTime.Time.Format(time.RFC3339)
	}

	return resp, nil
}
