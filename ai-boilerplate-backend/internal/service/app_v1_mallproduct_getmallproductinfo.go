package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// GetMallProductInfo 获取商品详情
func (a *AppV1MallProductService) GetMallProductInfo(ctx context.Context, req *pb.GetMallProductInfoReq) (*pb.GetMallProductInfoReply, error) {
	resp := &pb.GetMallProductInfoReply{}
	data, err := a.mallProductRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.MallProductInfo{
		Id:            data.ID,
		ProductType:   data.ProductType,
		ProductName:   data.ProductName,
		ProductDesc:   data.ProductDesc,
		ProductImages: string(data.ProductImages),
		ProductDetail: string(data.ProductDetail),
		ProductConfig: string(data.ProductConfig),
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
