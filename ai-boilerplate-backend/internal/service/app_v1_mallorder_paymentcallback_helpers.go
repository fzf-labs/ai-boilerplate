package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"gorm.io/datatypes"
)

const (
	paymentChannelWechat = "wechat"
	paymentMethodH5      = "h5"
	mockPaymentAppID     = "mock-wechat-app"
	mockPaymentSignType  = "HMAC-SHA256"
	mockPaymentSignKey   = "ai-boilerplate-mock-payment-sign-key"
)

type paymentCallbackInput struct {
	transactionID           string
	paymentChannel          string
	paymentMethod           string
	thirdPartyTransactionID string
	callbackJSON            datatypes.JSON
}

func buildPaymentCallbackInput(order *ai_boilerplate_model.MallOrder, req *pb.PaymentCallbackReq) paymentCallbackInput {
	transactionID := strings.TrimSpace(req.GetTransactionId())
	if transactionID == "" {
		transactionID = buildMockPaymentTransactionID(order.ID)
	}
	paymentChannel := strings.TrimSpace(req.GetPaymentMethod())
	if paymentChannel == "" {
		paymentChannel = order.PaymentMethod
	}
	if paymentChannel == "" {
		paymentChannel = paymentChannelWechat
	}
	return paymentCallbackInput{
		transactionID:           transactionID,
		paymentChannel:          paymentChannel,
		paymentMethod:           paymentMethodH5,
		thirdPartyTransactionID: strings.TrimSpace(req.GetTransactionId()),
		callbackJSON:            normalizePaymentCallbackData(req.GetCallbackData()),
	}
}

func normalizePaymentCallbackData(raw string) datatypes.JSON {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return datatypes.JSON([]byte("{}"))
	}
	if json.Valid([]byte(raw)) {
		return datatypes.JSON([]byte(raw))
	}
	encoded, _ := json.Marshal(map[string]string{"raw": raw})
	return datatypes.JSON(encoded)
}

func applyPaymentCallbackToOrder(order *ai_boilerplate_model.MallOrder, paymentStatus int32, paymentMethod string, now time.Time) {
	order.PaymentMethod = paymentMethod
	if paymentStatus == 1 {
		order.PaymentStatus = 1
		order.PaymentTime = sql.NullTime{Time: now, Valid: true}
		if order.ProductType == constant.MallProductTypeMembership.String() {
			order.Status = "completed"
			order.DeliveryTime = sql.NullTime{Time: now, Valid: true}
		} else {
			order.Status = "pendingDelivery"
		}
		return
	}
	order.PaymentStatus = 2
	order.Status = "canceled"
}

func applyPaymentCallbackToRecord(
	record *ai_boilerplate_model.MallPaymentRecord,
	order *ai_boilerplate_model.MallOrder,
	paymentStatus int32,
	input paymentCallbackInput,
	now time.Time,
	isNew bool,
) {
	if isNew {
		record.OrderID = order.ID
		record.TransactionID = input.transactionID
		record.ThirdPartyOrderNo = order.ID
		record.CreatedAt = now
		record.UpdatedAt = now
	}
	record.PaymentChannel = input.paymentChannel
	record.PaymentMethod = input.paymentMethod
	record.Amount = order.ActualAmount
	record.Currency = order.Currency
	record.Status = int32(constant.StatusEnable)
	record.PaymentStatus = paymentStatus
	record.ThirdPartyTransactionID = input.thirdPartyTransactionID
	record.CallbackData = input.callbackJSON
	record.CallbackTime = sql.NullTime{Time: now, Valid: true}
	if paymentStatus == 1 {
		record.ErrorMessage = ""
		return
	}
	record.ErrorMessage = "支付失败"
}

func isSupportedPaymentChannel(paymentMethod string) bool {
	return strings.TrimSpace(paymentMethod) == paymentChannelWechat
}

func buildMockPaymentTransactionID(orderID string) string {
	return "mock_" + strings.ReplaceAll(orderID, "-", "")
}

func buildMockPaymentURL(order *ai_boilerplate_model.MallOrder, transactionID string, timestamp string, nonce string, paySign string) string {
	values := url.Values{}
	values.Set("orderId", order.ID)
	values.Set("prepayId", transactionID)
	values.Set("amount", formatPaymentAmount(order.ActualAmount))
	values.Set("timestamp", timestamp)
	values.Set("nonce", nonce)
	values.Set("sign", paySign)
	return "https://mock-pay.local/wechat/checkout?" + values.Encode()
}

func buildMockPaymentSign(orderID string, transactionID string, paymentStatus int32, amount float64, timestamp string, nonce string) string {
	mac := hmac.New(sha256.New, []byte(mockPaymentSignKey))
	_, _ = mac.Write([]byte(mockPaymentSignContent(orderID, transactionID, paymentStatus, formatPaymentAmount(amount), timestamp, nonce)))
	return hex.EncodeToString(mac.Sum(nil))
}

func mockPaymentSignContent(orderID string, transactionID string, paymentStatus int32, amount string, timestamp string, nonce string) string {
	return strings.Join([]string{
		orderID,
		transactionID,
		strconv.FormatInt(int64(paymentStatus), 10),
		amount,
		timestamp,
		nonce,
	}, "\n")
}

func formatPaymentAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}

type mockPaymentCallbackData struct {
	OrderID       string `json:"orderId"`
	TransactionID string `json:"transactionId"`
	PaymentStatus int32  `json:"paymentStatus"`
	Amount        string `json:"amount"`
	Timestamp     string `json:"timestamp"`
	Nonce         string `json:"nonce"`
	Sign          string `json:"sign"`
}

func buildMockPaymentCallbackData(order *ai_boilerplate_model.MallOrder, transactionID string, paymentStatus int32, now time.Time) string {
	payload := mockPaymentCallbackData{
		OrderID:       order.ID,
		TransactionID: transactionID,
		PaymentStatus: paymentStatus,
		Amount:        formatPaymentAmount(order.ActualAmount),
		Timestamp:     strconv.FormatInt(now.Unix(), 10),
		Nonce:         strconv.FormatInt(now.UnixNano(), 36),
	}
	payload.Sign = buildMockPaymentSign(order.ID, transactionID, paymentStatus, order.ActualAmount, payload.Timestamp, payload.Nonce)
	data, _ := json.Marshal(payload)
	return string(data)
}

func verifyMockPaymentCallbackSignature(order *ai_boilerplate_model.MallOrder, req *pb.PaymentCallbackReq) error {
	raw := strings.TrimSpace(req.GetCallbackData())
	if raw == "" {
		return fmt.Errorf("callback data is empty")
	}
	var payload mockPaymentCallbackData
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		return fmt.Errorf("callback data is not valid json: %w", err)
	}
	transactionID := strings.TrimSpace(req.GetTransactionId())
	if transactionID == "" {
		return fmt.Errorf("transaction id is empty")
	}
	if payload.OrderID != order.ID {
		return fmt.Errorf("order id mismatch")
	}
	if payload.TransactionID != transactionID {
		return fmt.Errorf("transaction id mismatch")
	}
	if payload.PaymentStatus != req.GetPaymentStatus() {
		return fmt.Errorf("payment status mismatch")
	}
	if payload.Amount != formatPaymentAmount(order.ActualAmount) {
		return fmt.Errorf("amount mismatch")
	}
	if payload.Timestamp == "" || payload.Nonce == "" || payload.Sign == "" {
		return fmt.Errorf("signature payload is incomplete")
	}
	if _, err := strconv.ParseInt(payload.Timestamp, 10, 64); err != nil {
		return fmt.Errorf("timestamp is invalid: %w", err)
	}
	expected := buildMockPaymentSign(order.ID, transactionID, req.GetPaymentStatus(), order.ActualAmount, payload.Timestamp, payload.Nonce)
	if !hmac.Equal([]byte(payload.Sign), []byte(expected)) {
		return fmt.Errorf("signature mismatch")
	}
	return nil
}
