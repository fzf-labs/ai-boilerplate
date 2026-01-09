package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMallOrderInfo 订单信息表-单条数据查询
func (a *AdminV1MallOrderService) GetMallOrderInfo(ctx context.Context, req *pb.GetMallOrderInfoReq) (*pb.GetMallOrderInfoReply, error) {
	resp := &pb.GetMallOrderInfoReply{}
	data, err := a.mallOrderRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.MallOrderInfo{
		Id:             data.ID,
		UserId:         data.UserID,
		ProductType:    data.ProductType,
		ProductId:      data.ProductID,
		OriginalAmount: data.OriginalAmount,
		DiscountAmount: data.DiscountAmount,
		ActualAmount:   data.ActualAmount,
		RefundAmount:   data.RefundAmount,
		Currency:       data.Currency,
		PaymentMethod:  data.PaymentMethod,
		PaymentStatus:  data.PaymentStatus,
		PaymentTime:    timeutil.RFC3339(data.PaymentTime.Time),
		DeliveryTime:   timeutil.RFC3339(data.DeliveryTime.Time),
		ExpiredTime:    timeutil.RFC3339(data.ExpiredTime.Time),
		Remark:         data.Remark,
		Status:         data.Status,
		CreatedAt:      data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
