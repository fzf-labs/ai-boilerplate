package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetPaymentInfo 获取支付信息接口-选择支付方式并获取支付信息
func (a *AppV1MallOrderService) GetPaymentInfo(ctx context.Context, req *pb.GetPaymentInfoReq) (*pb.GetPaymentInfoReply, error) {
	resp := &pb.GetPaymentInfoReply{}

	// 1. 获取当前用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 2. 查询订单
	order, err := a.mallOrderRepo.FindOneCacheByID(ctx, req.GetOrderId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if order == nil || order.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	// 3. 校验订单归属
	if order.UserID != userID {
		return nil, pb.ErrorReasonUnauthorized()
	}

	// 4. 检查订单状态（必须是待支付状态）
	if order.PaymentStatus != 0 {
		return nil, pb.ErrorReasonParamError() // 订单已支付或已取消
	}

	// 5. 检查订单是否过期
	if order.ExpiredTime.Valid && time.Now().After(order.ExpiredTime.Time) {
		return nil, pb.ErrorReasonParamError() // 订单已过期
	}

	// 6. 更新订单支付方式
	if order.PaymentMethod != req.GetPaymentMethod() {
		oldData := a.mallOrderRepo.DeepCopy(order)
		order.PaymentMethod = req.GetPaymentMethod()
		if err := a.mallOrderRepo.UpdateOneCacheWithZero(ctx, order, oldData); err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}

	// 7. 构建支付信息响应
	resp.OrderId = order.ID
	resp.PaymentMethod = req.GetPaymentMethod()
	resp.ActualAmount = order.ActualAmount

	// 设置支付过期时间
	if order.ExpiredTime.Valid {
		resp.ExpiredTime = order.ExpiredTime.Time.Format(time.RFC3339)
	}

	// 8. 根据支付方式生成对应的支付信息
	// TODO: 这里应该调用对应的支付SDK生成真实的支付信息
	// 目前返回模拟数据，后续对接真实支付接口
	switch req.GetPaymentMethod() {
	case "wechat":
		// 微信支付信息（模拟）
		resp.AppId = "wx_app_id_placeholder"
		resp.PrepayId = fmt.Sprintf("wx_prepay_%s", order.ID)
		resp.TimeStamp = fmt.Sprintf("%d", time.Now().Unix())
		resp.NonceStr = fmt.Sprintf("nonce_%d", time.Now().UnixNano())
		resp.SignType = "RSA"
		resp.PaySign = "sign_placeholder"
		resp.QrCodeUrl = fmt.Sprintf("weixin://wxpay/bizpayurl?pr=%s", order.ID)
	case "alipay":
		// 支付宝支付信息（模拟）
		resp.PaymentUrl = fmt.Sprintf("https://openapi.alipay.com/gateway.do?order_id=%s", order.ID)
		resp.QrCodeUrl = fmt.Sprintf("https://qr.alipay.com/%s", order.ID)
	default:
		return nil, pb.ErrorReasonParamError() // 不支持的支付方式
	}

	return resp, nil
}
