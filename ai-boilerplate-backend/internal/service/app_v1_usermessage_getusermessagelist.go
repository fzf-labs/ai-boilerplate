package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserMessageList App-用户消息-列表数据查询
func (a *AppV1UserMessageService) GetUserMessageList(ctx context.Context, req *pb.GetUserMessageListReq) (*pb.GetUserMessageListReply, error) {
	resp := &pb.GetUserMessageListReply{
		Total: 0,
		List:  []*pb.UserMessageInfo{},
	}
	if req.GetCategory() == "" {
		return nil, pb.ErrorReasonParamError()
	}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "user_id",
				Value: userID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "category",
				Value: req.GetCategory(),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "sent_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetReadStatus() != 0 {
		query := &condition.QueryParam{
			Field: "read_at",
			Exp:   condition.ISNULL,
			Logic: condition.AND,
		}
		if req.GetReadStatus() > 0 {
			query.Exp = condition.ISNOTNULL
		}
		param.Query = append(param.Query, query)
	}
	list, p, err := a.userMessageRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = int32(p.Total)
	if len(list) > 0 {
		for _, v := range list {
			readAt := ""
			if v.ReadAt.Valid {
				readAt = timeutil.RFC3339(v.ReadAt.Time)
			}
			resp.List = append(resp.List, &pb.UserMessageInfo{
				Id:            v.ID,
				MessageId:     v.MessageID,
				UserId:        v.UserID,
				Category:      v.Category,
				Title:         v.Title,
				Summary:       v.Summary,
				CoverURL:      v.CoverURL,
				Content:       v.Content,
				LinkURL:       v.LinkURL,
				AudienceType:  v.AudienceType,
				AudienceValue: string(v.AudienceValue),
				SentAt:        timeutil.RFC3339(v.SentAt),
				ReadAt:        readAt,
				AdminId:       v.AdminID,
				CreatedAt:     timeutil.RFC3339(v.CreatedAt),
				UpdatedAt:     timeutil.RFC3339(v.UpdatedAt),
			})
		}
	}
	return resp, nil
}
