package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserMessageCategoryCounts App-用户消息-分类未读数量
func (a *AppV1UserMessageService) GetUserMessageCategoryCounts(ctx context.Context, req *pb.GetUserMessageCategoryCountsReq) (*pb.GetUserMessageCategoryCountsReply, error) {
	resp := &pb.GetUserMessageCategoryCountsReply{}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonParamError()
	}
	transactionUnread, err := a.countUnreadByCategory(ctx, userID, "transaction")
	if err != nil {
		return nil, err
	}
	systemUnread, err := a.countUnreadByCategory(ctx, userID, "system")
	if err != nil {
		return nil, err
	}
	serviceUnread, err := a.countUnreadByCategory(ctx, userID, "service")
	if err != nil {
		return nil, err
	}
	resp.TransactionUnread = transactionUnread
	resp.SystemUnread = systemUnread
	resp.ServiceUnread = serviceUnread
	resp.TotalUnread = transactionUnread + systemUnread + serviceUnread
	return resp, nil
}

func (a *AppV1UserMessageService) countUnreadByCategory(ctx context.Context, userID string, category string) (int32, error) {
	param := &condition.Req{
		Page:     1,
		PageSize: 1,
		Query: []*condition.QueryParam{
			{
				Field: "user_id",
				Value: userID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "category",
				Value: category,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "read_at",
				Exp:   condition.ISNULL,
				Logic: condition.AND,
			},
		},
	}
	_, p, err := a.userMessageRepo.FindMultiByCondition(ctx, param)
	if err != nil {
		return 0, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return int32(p.Total), nil
}
