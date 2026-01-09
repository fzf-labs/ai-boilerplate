package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxXcxUser 小程序用户表-删除一条数据
func (a *AdminV1WxXcxUserService) DeleteWxXcxUser(ctx context.Context, req *pb.DeleteWxXcxUserReq) (*pb.DeleteWxXcxUserReply, error) {
	resp := &pb.DeleteWxXcxUserReply{}
	err := a.wxXcxUserRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
