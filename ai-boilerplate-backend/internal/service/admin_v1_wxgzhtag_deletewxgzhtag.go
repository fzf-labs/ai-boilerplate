package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxGzhTag 公众号标签表-删除一条数据
func (a *AdminV1WxGzhTagService) DeleteWxGzhTag(ctx context.Context, req *pb.DeleteWxGzhTagReq) (*pb.DeleteWxGzhTagReply, error) {
	resp := &pb.DeleteWxGzhTagReply{}
	err := a.wxGzhTagRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
