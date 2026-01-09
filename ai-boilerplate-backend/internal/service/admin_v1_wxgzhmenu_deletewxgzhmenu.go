package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxGzhMenu 公众号菜单表-删除一条数据
func (a *AdminV1WxGzhMenuService) DeleteWxGzhMenu(ctx context.Context, req *pb.DeleteWxGzhMenuReq) (*pb.DeleteWxGzhMenuReply, error) {
	resp := &pb.DeleteWxGzhMenuReply{}
	err := a.wxGzhMenuRepo.DeleteOneCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
