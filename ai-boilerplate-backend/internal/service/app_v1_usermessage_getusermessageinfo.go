package service

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserMessageInfo App-用户消息-单条数据查询
func (a *AppV1UserMessageService) GetUserMessageInfo(ctx context.Context, req *pb.GetUserMessageInfoReq) (*pb.GetUserMessageInfoReply, error) {
	resp := &pb.GetUserMessageInfoReply{}
	data, err := a.userMessageRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if data.UserID != userID {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if !data.ReadAt.Valid {
		oldData := a.userMessageRepo.DeepCopy(data)
		data.ReadAt = sql.NullTime{Time: time.Now(), Valid: true}
		if err := a.userMessageRepo.UpdateOneCacheWithZero(ctx, data, oldData); err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
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
