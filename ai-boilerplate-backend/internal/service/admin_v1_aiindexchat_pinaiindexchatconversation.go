package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// PinAiIndexChatConversation AI 聊天对话表-置顶和取消置顶
func (a *AdminV1AiIndexChatService) PinAiIndexChatConversation(ctx context.Context, req *pb.PinAiIndexChatConversationReq) (*pb.PinAiIndexChatConversationReply, error) {
	resp := &pb.PinAiIndexChatConversationReply{}
	data, err := a.aiChatConversationRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiChatConversationRepo.DeepCopy(data)
	data.Pinned = req.GetPinned()
	if req.GetPinned() {
		data.PinnedTime = timeutil.TimeToSQLNullTime(carbon.Parse(time.Now().Format(time.RFC3339)).StdTime())
	} else {
		data.PinnedTime = sql.NullTime{}
	}
	err = a.aiChatConversationRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
