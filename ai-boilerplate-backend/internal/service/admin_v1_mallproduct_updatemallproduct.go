package service

import (
	"context"
	"encoding/json"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMallProduct 商品表-更新一条数据
func (a *AdminV1MallProductService) UpdateMallProduct(ctx context.Context, req *pb.UpdateMallProductReq) (*pb.UpdateMallProductReply, error) {
	resp := &pb.UpdateMallProductReply{}
	data, err := a.mallProductRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	productImages, err := json.Marshal(req.GetProductImages())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	productDetail, err := json.Marshal(req.GetProductDetail())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	productConfig, err := json.Marshal(req.GetProductConfig())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	oldData := a.mallProductRepo.DeepCopy(data)
	data.ProductType = req.GetProductType()
	data.ProductName = req.GetProductName()
	data.ProductDesc = req.GetProductDesc()
	data.ProductImages = productImages
	data.ProductDetail = productDetail
	data.ProductConfig = productConfig
	data.OriginalPrice = req.GetOriginalPrice()
	data.CurrentPrice = req.GetCurrentPrice()
	data.StockQuantity = req.GetStockQuantity()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.mallProductRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
