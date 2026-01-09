package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateWxGzhAutoReplyStatus 公众号消息自动回复表-更新状态
func (a *AdminV1WxGzhAutoReplyService) UpdateWxGzhAutoReplyStatus(ctx context.Context, req *pb.UpdateWxGzhAutoReplyStatusReq) (*pb.UpdateWxGzhAutoReplyStatusReply, error) {
	resp := &pb.UpdateWxGzhAutoReplyStatusReply{}
	data, err := a.wxGzhAutoReplyRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.wxGzhAutoReplyRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.wxGzhAutoReplyRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
