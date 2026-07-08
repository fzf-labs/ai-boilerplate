package service

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"gorm.io/datatypes"
)

type paymentCallbackInput struct {
	transactionID           string
	paymentMethod           string
	thirdPartyTransactionID string
	callbackJSON            datatypes.JSON
}

func buildPaymentCallbackInput(order *ai_boilerplate_model.MallOrder, req *pb.PaymentCallbackReq) paymentCallbackInput {
	transactionID := strings.TrimSpace(req.GetTransactionId())
	if transactionID == "" {
		transactionID = order.ID
	}
	paymentMethod := strings.TrimSpace(req.GetPaymentMethod())
	if paymentMethod == "" {
		paymentMethod = order.PaymentMethod
	}
	return paymentCallbackInput{
		transactionID:           transactionID,
		paymentMethod:           paymentMethod,
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
	record.PaymentChannel = input.paymentMethod
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
