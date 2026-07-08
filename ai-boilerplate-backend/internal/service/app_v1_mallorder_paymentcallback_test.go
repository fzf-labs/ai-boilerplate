package service

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/testfixture"
)

func TestBuildPaymentCallbackInputUsesFallbacksAndNormalizesRawCallback(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeMembership.String())
	req := testfixture.PaymentCallbackReq(order.ID, 1)
	req.TransactionId = "  "
	req.PaymentMethod = " "
	req.CallbackData = "provider=ok"

	got := buildPaymentCallbackInput(order, req)
	wantTransactionID := buildMockPaymentTransactionID(order.ID)
	if got.transactionID != wantTransactionID {
		t.Fatalf("transactionID = %q, want %q", got.transactionID, wantTransactionID)
	}
	if got.paymentChannel != order.PaymentMethod {
		t.Fatalf("paymentChannel = %q, want %q", got.paymentChannel, order.PaymentMethod)
	}
	if got.paymentMethod != paymentMethodH5 {
		t.Fatalf("paymentMethod = %q, want %q", got.paymentMethod, paymentMethodH5)
	}
	if got.thirdPartyTransactionID != "" {
		t.Fatalf("thirdPartyTransactionID = %q, want empty trimmed value", got.thirdPartyTransactionID)
	}

	var callback map[string]string
	if err := json.Unmarshal(got.callbackJSON, &callback); err != nil {
		t.Fatalf("callback json should be valid: %v", err)
	}
	if callback["raw"] != "provider=ok" {
		t.Fatalf("callback raw = %q", callback["raw"])
	}
}

func TestNormalizePaymentCallbackDataKeepsValidJSON(t *testing.T) {
	t.Parallel()

	got := normalizePaymentCallbackData(` {"paid":true} `)
	var callback map[string]bool
	if err := json.Unmarshal(got, &callback); err != nil {
		t.Fatalf("callback json should be valid: %v", err)
	}
	if !callback["paid"] {
		t.Fatalf("paid flag missing from callback: %#v", callback)
	}
}

func TestApplyPaymentCallbackToMembershipOrderSuccess(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeMembership.String())
	now := time.Date(2026, 2, 3, 4, 5, 6, 0, time.UTC)

	applyPaymentCallbackToOrder(order, 1, "alipay", now)

	if order.PaymentMethod != "alipay" || order.PaymentStatus != 1 || order.Status != "completed" {
		t.Fatalf("unexpected paid membership order: %#v", order)
	}
	if !order.PaymentTime.Valid || !order.PaymentTime.Time.Equal(now) {
		t.Fatalf("payment time = %#v, want %s", order.PaymentTime, now)
	}
	if !order.DeliveryTime.Valid || !order.DeliveryTime.Time.Equal(now) {
		t.Fatalf("membership delivery time = %#v, want %s", order.DeliveryTime, now)
	}
}

func TestApplyPaymentCallbackToServiceOrderSuccess(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeService.String())
	now := time.Date(2026, 2, 3, 4, 5, 6, 0, time.UTC)

	applyPaymentCallbackToOrder(order, 1, "wechat", now)

	if order.PaymentStatus != 1 || order.Status != "pendingDelivery" {
		t.Fatalf("unexpected paid service order: %#v", order)
	}
	if order.DeliveryTime.Valid {
		t.Fatalf("non-membership delivery time should not be set: %#v", order.DeliveryTime)
	}
}

func TestApplyPaymentCallbackToOrderFailure(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeGoods.String())

	applyPaymentCallbackToOrder(order, 2, "wechat", time.Now())

	if order.PaymentStatus != 2 || order.Status != "canceled" {
		t.Fatalf("unexpected failed order: %#v", order)
	}
	if order.PaymentTime.Valid || order.DeliveryTime.Valid {
		t.Fatalf("failed order should not set payment or delivery time: %#v %#v", order.PaymentTime, order.DeliveryTime)
	}
}

func TestApplyPaymentCallbackToRecord(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeMembership.String())
	req := testfixture.PaymentCallbackReq(order.ID, 1)
	input := buildPaymentCallbackInput(order, req)
	record := &ai_boilerplate_model.MallPaymentRecord{}
	now := time.Date(2026, 2, 3, 4, 5, 6, 0, time.UTC)

	applyPaymentCallbackToRecord(record, order, 1, input, now, true)

	if record.OrderID != order.ID || record.TransactionID != req.TransactionId {
		t.Fatalf("unexpected record identity: %#v", record)
	}
	if record.PaymentChannel != paymentChannelWechat || record.PaymentMethod != paymentMethodH5 {
		t.Fatalf("unexpected record payment route: %#v", record)
	}
	if record.Amount != order.ActualAmount || record.Currency != order.Currency {
		t.Fatalf("unexpected record amount: %#v", record)
	}
	if record.PaymentStatus != 1 || record.ErrorMessage != "" {
		t.Fatalf("unexpected successful record status: %#v", record)
	}
	if !record.CallbackTime.Valid || !record.CallbackTime.Time.Equal(now) {
		t.Fatalf("callback time = %#v, want %s", record.CallbackTime, now)
	}

	applyPaymentCallbackToRecord(record, order, 2, input, now, false)
	if record.PaymentStatus != 2 || record.ErrorMessage != "支付失败" {
		t.Fatalf("unexpected failed record status: %#v", record)
	}
}

func TestMockPaymentCallbackSignature(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeMembership.String())
	transactionID := buildMockPaymentTransactionID(order.ID)
	now := time.Date(2026, 2, 3, 4, 5, 6, 0, time.UTC)
	req := testfixture.PaymentCallbackReq(order.ID, 1)
	req.TransactionId = transactionID
	req.CallbackData = buildMockPaymentCallbackData(order, transactionID, 1, now)

	if err := verifyMockPaymentCallbackSignature(order, req); err != nil {
		t.Fatalf("signature should verify: %v", err)
	}

	req.PaymentStatus = 2
	if err := verifyMockPaymentCallbackSignature(order, req); err == nil {
		t.Fatalf("signature should fail after status tampering")
	}
}

func TestBuildMockPaymentURLIncludesSignedPrepayPayload(t *testing.T) {
	t.Parallel()

	order := testfixture.MallOrder(constant.MallProductTypeMembership.String())
	transactionID := buildMockPaymentTransactionID(order.ID)
	timestamp := "1770110706"
	nonce := "fixture-nonce"
	sign := buildMockPaymentSign(order.ID, transactionID, 0, order.ActualAmount, timestamp, nonce)

	got := buildMockPaymentURL(order, transactionID, timestamp, nonce, sign)
	if !strings.Contains(got, "https://mock-pay.local/wechat/checkout?") {
		t.Fatalf("unexpected mock payment url: %s", got)
	}
	if !strings.Contains(got, "prepayId="+transactionID) || !strings.Contains(got, "sign="+sign) {
		t.Fatalf("mock payment url missing signed payload: %s", got)
	}
}
