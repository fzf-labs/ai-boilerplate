package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/samber/lo"
)

// GetMallPaymentRecordSuccessByOrderId 支付记录表-单条数据查询
//
//nolint:revive // gRPC method name comes from proto.
func (a *AdminV1MallPaymentRecordService) GetMallPaymentRecordSuccessByOrderId(ctx context.Context, req *pb.GetMallPaymentRecordSuccessByOrderIdReq) (*pb.GetMallPaymentRecordSuccessByOrderIdReply, error) {
	resp := &pb.GetMallPaymentRecordSuccessByOrderIdReply{
		Info: &pb.MallPaymentRecordInfo{},
	}
	list, err := a.mallPaymentRecordRepo.FindMultiCacheByOrderID(ctx, req.GetOrderId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	list = lo.Filter(list, func(item *ai_boilerplate_model.MallPaymentRecord, _ int) bool {
		return item.PaymentStatus == 1
	})
	if len(list) > 0 {
		resp.Info = &pb.MallPaymentRecordInfo{
			Id:                      list[0].ID,
			OrderId:                 list[0].OrderID,
			TransactionId:           list[0].TransactionID,
			PaymentChannel:          list[0].PaymentChannel,
			PaymentMethod:           list[0].PaymentMethod,
			Amount:                  list[0].Amount,
			Currency:                list[0].Currency,
			PaymentStatus:           list[0].PaymentStatus,
			ThirdPartyOrderNo:       list[0].ThirdPartyOrderNo,
			ThirdPartyTransactionId: list[0].ThirdPartyTransactionID,
			CallbackData:            string(list[0].CallbackData),
			CallbackTime:            list[0].CallbackTime.Time.Format(time.RFC3339),
			ErrorCode:               list[0].ErrorCode,
			ErrorMessage:            list[0].ErrorMessage,
			Status:                  list[0].Status,
			CreatedAt:               list[0].CreatedAt.Format(time.RFC3339),
			UpdatedAt:               list[0].UpdatedAt.Format(time.RFC3339),
		}
	}
	return resp, nil
}
