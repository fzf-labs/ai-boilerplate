package service

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateUserMessageRead App-用户消息-标记已读
func (a *AppV1UserMessageService) UpdateUserMessageRead(ctx context.Context, req *pb.UpdateUserMessageReadReq) (*pb.UpdateUserMessageReadReply, error) {
	resp := &pb.UpdateUserMessageReadReply{}
	data, err := a.userMessageRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if data.UserID != userID {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if data.ReadAt.Valid {
		return resp, nil
	}
	oldData := a.userMessageRepo.DeepCopy(data)
	data.ReadAt = sql.NullTime{Time: time.Now(), Valid: true}
	if err := a.userMessageRepo.UpdateOneCacheWithZero(ctx, data, oldData); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
