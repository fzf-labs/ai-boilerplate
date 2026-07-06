package service

import (
	"fmt"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
)

func buildPaymentInfoReply(order *ai_boilerplate_model.MallOrder, paymentMethod string, now time.Time) (*pb.GetPaymentInfoReply, error) {
	resp := &pb.GetPaymentInfoReply{
		OrderId:       order.ID,
		PaymentMethod: paymentMethod,
		ActualAmount:  order.ActualAmount,
	}
	if order.ExpiredTime.Valid {
		resp.ExpiredTime = order.ExpiredTime.Time.Format(time.RFC3339)
	}

	switch paymentMethod {
	case "wechat":
		resp.AppId = "wx_app_id_placeholder"
		resp.PrepayId = fmt.Sprintf("wx_prepay_%s", order.ID)
		resp.TimeStamp = fmt.Sprintf("%d", now.Unix())
		resp.NonceStr = fmt.Sprintf("nonce_%d", now.UnixNano())
		resp.SignType = "RSA"
		resp.PaySign = "sign_placeholder"
		resp.QrCodeUrl = fmt.Sprintf("weixin://wxpay/bizpayurl?pr=%s", order.ID)
	case "alipay":
		resp.PaymentUrl = fmt.Sprintf("https://openapi.alipay.com/gateway.do?order_id=%s", order.ID)
		resp.QrCodeUrl = fmt.Sprintf("https://qr.alipay.com/%s", order.ID)
	default:
		return nil, pb.ErrorReasonParamError()
	}

	return resp, nil
}
