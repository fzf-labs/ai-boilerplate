package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserOrderList 查询订单接口-用户订单列表
func (a *AppV1MallOrderService) GetUserOrderList(ctx context.Context, req *pb.GetUserOrderListReq) (*pb.GetUserOrderListReply, error) {
	resp := &pb.GetUserOrderListReply{
		Total: 0,
		List:  []*pb.MallOrderInfo{},
	}

	// 1. 获取当前用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 2. 构建查询条件
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "user_id",
				Value: userID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{Field: "created_at", Order: condition.DESC},
		},
	}

	// 3. 添加订单状态过滤
	if req.GetStatus() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}

	// 4. 添加支付状态过滤
	if req.GetPaymentStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "payment_status",
			Value: req.GetPaymentStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}

	// 5. 查询订单列表
	list, p, err := a.mallOrderRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total

	// 6. 构建返回列表
	for _, order := range list {
		info := &pb.MallOrderInfo{
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
			info.PaymentTime = order.PaymentTime.Time.Format(time.RFC3339)
		}
		if order.DeliveryTime.Valid {
			info.DeliveryTime = order.DeliveryTime.Time.Format(time.RFC3339)
		}
		if order.ExpiredTime.Valid {
			info.ExpiredTime = order.ExpiredTime.Time.Format(time.RFC3339)
		}

		resp.List = append(resp.List, info)
	}

	return resp, nil
}
