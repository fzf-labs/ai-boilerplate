package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetUserMessageInfo App-用户消息-单条数据查询
func (a *AdminV1UserMessageService) GetUserMessageInfo(ctx context.Context, req *pb.GetUserMessageInfoReq) (*pb.GetUserMessageInfoReply, error) {
	resp := &pb.GetUserMessageInfoReply{}
	data, err := a.userMessageRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	readAt := ""
	if data.ReadAt.Valid {
		readAt = timeutil.RFC3339(data.ReadAt.Time)
	}
	resp.Info = &pb.UserMessageInfo{
		Id:            data.ID,
		MessageId:     data.MessageID,
		UserId:        data.UserID,
		Category:      data.Category,
		Title:         data.Title,
		Summary:       data.Summary,
		CoverURL:      data.CoverURL,
		Content:       data.Content,
		LinkURL:       data.LinkURL,
		AudienceType:  data.AudienceType,
		AudienceValue: string(data.AudienceValue),
		SentAt:        timeutil.RFC3339(data.SentAt),
		ReadAt:        readAt,
		AdminId:       data.AdminID,
		CreatedAt:     timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:     timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
