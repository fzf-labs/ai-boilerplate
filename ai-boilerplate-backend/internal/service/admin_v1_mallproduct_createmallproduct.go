package service

import (
	"context"
	"encoding/json"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateMallProduct 商品表-创建一条数据
func (a *AdminV1MallProductService) CreateMallProduct(ctx context.Context, req *pb.CreateMallProductReq) (*pb.CreateMallProductReply, error) {
	resp := &pb.CreateMallProductReply{}
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
	data := a.mallProductRepo.NewData()
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
	err = a.mallProductRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
