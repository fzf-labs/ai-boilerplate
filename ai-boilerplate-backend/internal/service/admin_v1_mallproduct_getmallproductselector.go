package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMallProductSelector 商品表-选择器
func (a *AdminV1MallProductService) GetMallProductSelector(ctx context.Context, req *pb.GetMallProductSelectorReq) (*pb.GetMallProductSelectorReply, error) {
	resp := &pb.GetMallProductSelectorReply{
		List: []*pb.MallProductSelector{},
	}
	param := &condition.Req{
		Query: []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetSearchName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "product_name",
			Value: "%" + req.GetSearchName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	list, _, err := a.mallProductRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MallProductSelector{
				Id:          v.ID,
				ProductName: v.ProductName,
			})
		}
	}
	return resp, nil
}
