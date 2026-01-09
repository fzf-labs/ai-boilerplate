package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMallOrderList 订单信息表-列表数据查询
func (a *AdminV1MallOrderService) GetMallOrderList(ctx context.Context, req *pb.GetMallOrderListReq) (*pb.GetMallOrderListReply, error) {
	resp := &pb.GetMallOrderListReply{
		Total: 0,
		List:  []*pb.MallOrderInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.mallOrderRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MallOrderInfo{
				Id:             v.ID,
				UserId:         v.UserID,
				ProductType:    v.ProductType,
				ProductId:      v.ProductID,
				OriginalAmount: v.OriginalAmount,
				DiscountAmount: v.DiscountAmount,
				ActualAmount:   v.ActualAmount,
				RefundAmount:   v.RefundAmount,
				Currency:       v.Currency,
				PaymentMethod:  v.PaymentMethod,
				PaymentStatus:  v.PaymentStatus,
				PaymentTime:    timeutil.RFC3339(v.PaymentTime.Time),
				DeliveryTime:   timeutil.RFC3339(v.DeliveryTime.Time),
				ExpiredTime:    timeutil.RFC3339(v.ExpiredTime.Time),
				Remark:         v.Remark,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
