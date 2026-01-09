package service

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMallProductList 商品表-列表数据查询
func (a *AdminV1MallProductService) GetMallProductList(ctx context.Context, req *pb.GetMallProductListReq) (*pb.GetMallProductListReply, error) {
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
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetProductType() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "product_type",
			Value: req.GetProductType(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetProductName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "product_name",
			Value: req.GetProductName(),
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
			productImages := make([]string, 0)
			if v.ProductImages.String() != "" {
				err = json.Unmarshal(v.ProductImages, &productImages)
				if err != nil {
					return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
				}
			}
			productDetail := make([]string, 0)
			if v.ProductDetail.String() != "" {
				err = json.Unmarshal(v.ProductDetail, &productDetail)
				if err != nil {
					return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
				}
			}
			productConfig := &pb.ProductConfig{}
			if v.ProductConfig.String() != "" {
				err = json.Unmarshal(v.ProductConfig, productConfig)
				if err != nil {
					return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
				}
			}
			resp.List = append(resp.List, &pb.MallProductInfo{
				Id:            v.ID,
				ProductType:   v.ProductType,
				ProductName:   v.ProductName,
				ProductDesc:   v.ProductDesc,
				ProductImages: productImages,
				ProductDetail: productDetail,
				ProductConfig: productConfig,
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
