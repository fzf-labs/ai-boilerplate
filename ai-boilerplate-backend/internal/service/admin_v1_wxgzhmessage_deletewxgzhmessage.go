package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxGzhMessage 公众号消息表 -删除一条数据
func (a *AdminV1WxGzhMessageService) DeleteWxGzhMessage(ctx context.Context, req *pb.DeleteWxGzhMessageReq) (*pb.DeleteWxGzhMessageReply, error) {
	resp := &pb.DeleteWxGzhMessageReply{}
	err := a.wxGzhMessageRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
