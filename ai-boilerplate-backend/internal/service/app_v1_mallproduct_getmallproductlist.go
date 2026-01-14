package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMallProductList 获取商品列表
func (a *AppV1MallProductService) GetMallProductList(ctx context.Context, req *pb.GetMallProductListReq) (*pb.GetMallProductListReply, error) {
	resp := &pb.GetMallProductListReply{
		Total: 0,
		List:  []*pb.MallProductInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "sort",
				Order: condition.ASC,
			},
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}

	// 添加商品类型过滤
	if req.GetProductType() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "product_type",
			Value: req.GetProductType(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}

	// 添加状态过滤，默认只显示在售商品(status=1)
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	} else {
		// 默认只显示在售商品
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: int32(1),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}

	list, p, err := a.mallProductRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MallProductInfo{
				Id:            v.ID,
				ProductType:   v.ProductType,
				ProductName:   v.ProductName,
				ProductDesc:   v.ProductDesc,
				ProductImages: string(v.ProductImages),
				ProductDetail: string(v.ProductDetail),
				ProductConfig: string(v.ProductConfig),
				OriginalPrice: v.OriginalPrice,
				CurrentPrice:  v.CurrentPrice,
				StockQuantity: v.StockQuantity,
				SoldQuantity:  v.SoldQuantity,
				Sort:          v.Sort,
				Status:        v.Status,
				CreatedAt:     v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:     v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
