package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	adminpb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	apppb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
)

func TestSmsChannelInfoFromModelMasksSecret(t *testing.T) {
	now := time.Date(2026, 7, 6, 12, 0, 0, 0, time.UTC)
	info := smsChannelInfoFromModel(&ai_boilerplate_model.SmsChannel{
		ID:          "sms-channel-fixture",
		Name:        "Fixture SMS",
		Operator:    constant.SmsChannelCodeALIYUN.String(),
		APIKey:      "api-key-fixture",
		APISecret:   "secret-fixture-value",
		CallbackURL: "https://example.test/sms/callback",
		Status:      1,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	if info.APISecret == "secret-fixture-value" {
		t.Fatal("sms channel APISecret should be masked")
	}
	if info.APISecret != "secr****alue" {
		t.Fatalf("masked APISecret = %q, want secr****alue", info.APISecret)
	}
	if info.APIKey != "api-key-fixture" {
		t.Fatalf("APIKey = %q, want raw non-secret API key", info.APIKey)
	}
	if info.OperatorName == "" {
		t.Fatal("OperatorName should be populated from operator code")
	}
}

func TestMaskSecretHandlesShortValues(t *testing.T) {
	tests := []struct {
		name   string
		secret string
		want   string
	}{
		{name: "empty", secret: "", want: ""},
		{name: "short", secret: "abc", want: "****"},
		{name: "middle", secret: "abcdef", want: "ab****ef"},
		{name: "long", secret: "abcdefghijkl", want: "abcd****ijkl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskSecret(tt.secret); got != tt.want {
				t.Fatalf("maskSecret(%q) = %q, want %q", tt.secret, got, tt.want)
			}
		})
	}
}

func TestBuildPaymentInfoReplySupportsWechatAndAlipay(t *testing.T) {
	expiredAt := time.Date(2026, 7, 6, 12, 30, 0, 0, time.UTC)
	now := time.Unix(1_788_000_000, 123)
	order := &ai_boilerplate_model.MallOrder{
		ID:           "order-fixture",
		ActualAmount: 99.5,
		ExpiredTime:  sql.NullTime{Time: expiredAt, Valid: true},
	}

	wechat, err := buildPaymentInfoReply(order, "wechat", now)
	if err != nil {
		t.Fatalf("buildPaymentInfoReply(wechat) error = %v", err)
	}
	if wechat.OrderId != order.ID || wechat.PaymentMethod != "wechat" || wechat.ActualAmount != order.ActualAmount {
		t.Fatalf("wechat reply core fields = %+v", wechat)
	}
	if wechat.PrepayId != "wx_prepay_order-fixture" || wechat.TimeStamp != "1788000000" || wechat.NonceStr != "nonce_1788000000000000123" {
		t.Fatalf("wechat payment fixture fields = %+v", wechat)
	}
	if wechat.QrCodeUrl == "" || wechat.PaySign == "" || wechat.ExpiredTime != expiredAt.Format(time.RFC3339) {
		t.Fatalf("wechat payment reply missing expected fields: %+v", wechat)
	}

	alipay, err := buildPaymentInfoReply(order, "alipay", now)
	if err != nil {
		t.Fatalf("buildPaymentInfoReply(alipay) error = %v", err)
	}
	if alipay.PaymentUrl == "" || alipay.QrCodeUrl == "" || alipay.PaymentMethod != "alipay" {
		t.Fatalf("alipay payment reply missing expected fields: %+v", alipay)
	}
}

func TestBuildPaymentInfoReplyRejectsUnsupportedPaymentMethod(t *testing.T) {
	_, err := buildPaymentInfoReply(&ai_boilerplate_model.MallOrder{ID: "order-fixture"}, "wire-transfer", time.Now())
	if !apppb.IsParamError(err) {
		t.Fatalf("buildPaymentInfoReply() error = %v, want ParamError", err)
	}
}

func TestMallOrderServiceRequiresAuthenticatedUserBeforeRepos(t *testing.T) {
	svc := &AppV1MallOrderService{}

	if _, err := svc.PlaceOrder(context.Background(), &apppb.PlaceOrderReq{ProductId: "product-fixture"}); !apppb.IsUnauthorized(err) {
		t.Fatalf("PlaceOrder() error = %v, want Unauthorized", err)
	}
	if _, err := svc.GetPaymentInfo(context.Background(), &apppb.GetPaymentInfoReq{OrderId: "order-fixture", PaymentMethod: "wechat"}); !apppb.IsUnauthorized(err) {
		t.Fatalf("GetPaymentInfo() error = %v, want Unauthorized", err)
	}
}

func TestCurrentNoopUserAndLogoutContracts(t *testing.T) {
	userSvc := &AppV1UserService{}
	if reply, err := userSvc.SendVerifyCode(context.Background(), &apppb.SendVerifyCodeReq{Phone: "13800138000"}); err != nil || reply == nil {
		t.Fatalf("SendVerifyCode() reply = %v error = %v, want empty reply without error", reply, err)
	}
	if reply, err := userSvc.DeleteAccount(context.Background(), &apppb.DeleteAccountReq{Password: "fixture-password"}); err != nil || reply == nil {
		t.Fatalf("DeleteAccount() reply = %v error = %v, want empty reply without error", reply, err)
	}

	authSvc := &AdminV1SysAuthService{}
	if reply, err := authSvc.SysAuthLogout(context.Background(), &adminpb.SysAuthLogoutReq{}); err != nil || reply == nil {
		t.Fatalf("SysAuthLogout() reply = %v error = %v, want empty reply without error", reply, err)
	}
}
