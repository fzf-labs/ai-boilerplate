package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxGzhAutoReply 公众号消息自动回复表-删除一条数据
func (a *AdminV1WxGzhAutoReplyService) DeleteWxGzhAutoReply(ctx context.Context, req *pb.DeleteWxGzhAutoReplyReq) (*pb.DeleteWxGzhAutoReplyReply, error) {
	resp := &pb.DeleteWxGzhAutoReplyReply{}
	err := a.wxGzhAutoReplyRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
