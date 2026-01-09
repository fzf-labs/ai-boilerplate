package service

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetMallProductInfo 商品表-单条数据查询
func (a *AdminV1MallProductService) GetMallProductInfo(ctx context.Context, req *pb.GetMallProductInfoReq) (*pb.GetMallProductInfoReply, error) {
	resp := &pb.GetMallProductInfoReply{}
	data, err := a.mallProductRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	productImages := make([]string, 0)
	if data.ProductImages.String() != "" {
		err = json.Unmarshal(data.ProductImages, &productImages)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	productDetail := make([]string, 0)
	if data.ProductDetail.String() != "" {
		err = json.Unmarshal(data.ProductDetail, &productDetail)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	productConfig := &pb.ProductConfig{}
	if data.ProductConfig.String() != "" {
		err = json.Unmarshal(data.ProductConfig, productConfig)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	resp.Info = &pb.MallProductInfo{
		Id:            data.ID,
		ProductType:   data.ProductType,
		ProductName:   data.ProductName,
		ProductDesc:   data.ProductDesc,
		ProductImages: productImages,
		ProductDetail: productDetail,
		ProductConfig: productConfig,
		OriginalPrice: data.OriginalPrice,
		CurrentPrice:  data.CurrentPrice,
		StockQuantity: data.StockQuantity,
		SoldQuantity:  data.SoldQuantity,
		Sort:          data.Sort,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
