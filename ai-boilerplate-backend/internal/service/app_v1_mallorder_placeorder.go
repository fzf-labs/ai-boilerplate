package service

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// PlaceOrder 下单接口-创建订单
func (a *AppV1MallOrderService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderReq) (*pb.PlaceOrderReply, error) {
	resp := &pb.PlaceOrderReply{}

	// 1. 获取当前用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 2. 查询商品信息
	product, err := a.mallProductRepo.FindOneCacheByID(ctx, req.GetProductId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if product == nil || product.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	// 3. 校验商品状态和库存
	if product.Status != 1 {
		return nil, pb.ErrorReasonParamError()
	}
	if product.StockQuantity <= 0 {
		return nil, pb.ErrorReasonParamError()
	}

	// 4. 计算订单金额
	originalAmount := product.OriginalPrice
	actualAmount := product.CurrentPrice
	discountAmount := originalAmount - actualAmount

	// 5. 创建订单
	data := a.mallOrderRepo.NewData()
	data.UserID = userID
	data.ProductType = req.GetProductType()
	data.ProductID = req.GetProductId()
	data.OriginalAmount = originalAmount
	data.DiscountAmount = discountAmount
	data.ActualAmount = actualAmount
	data.RefundAmount = 0
	data.Currency = "CNY"
	data.PaymentMethod = req.GetPaymentMethod()
	data.PaymentStatus = 0 // 待支付
	data.Status = "pendingPayment"
	data.Remark = req.GetRemark()

	// 设置订单过期时间（30分钟后）
	expiredTime := time.Now().Add(30 * time.Minute)
	data.ExpiredTime = sql.NullTime{Time: expiredTime, Valid: true}

	if err := a.mallOrderRepo.CreateOneCache(ctx, data); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 6. 返回订单信息
	resp.OrderId = data.ID
	resp.ActualAmount = data.ActualAmount
	resp.PaymentInfo = "支付信息待实现（预留支付接口）"

	return resp, nil
}
