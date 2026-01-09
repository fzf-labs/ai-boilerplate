package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetMallPaymentRecordListByOrderId 支付记录表-列表数据查询
//
//nolint:revive // gRPC method name comes from proto.
func (a *AdminV1MallPaymentRecordService) GetMallPaymentRecordListByOrderId(ctx context.Context, req *pb.GetMallPaymentRecordListByOrderIdReq) (*pb.GetMallPaymentRecordListByOrderIdReply, error) {
	resp := &pb.GetMallPaymentRecordListByOrderIdReply{
		List: []*pb.MallPaymentRecordInfo{},
	}
	list, err := a.mallPaymentRecordRepo.FindMultiCacheByOrderID(ctx, req.GetOrderId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MallPaymentRecordInfo{
				Id:                      v.ID,
				OrderId:                 v.OrderID,
				TransactionId:           v.TransactionID,
				PaymentChannel:          v.PaymentChannel,
				PaymentMethod:           v.PaymentMethod,
				Amount:                  v.Amount,
				Currency:                v.Currency,
				PaymentStatus:           v.PaymentStatus,
				ThirdPartyOrderNo:       v.ThirdPartyOrderNo,
				ThirdPartyTransactionId: v.ThirdPartyTransactionID,
				CallbackData:            string(v.CallbackData),
				CallbackTime:            v.CallbackTime.Time.Format(time.RFC3339),
				ErrorCode:               v.ErrorCode,
				ErrorMessage:            v.ErrorMessage,
				Status:                  v.Status,
				CreatedAt:               v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:               v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
