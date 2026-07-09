package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhAccountInfo 公众号账号表-单条数据查询
func (a *AdminV1WxGzhAccountService) GetWxGzhAccountInfo(ctx context.Context, req *pb.GetWxGzhAccountInfoReq) (*pb.GetWxGzhAccountInfoReply, error) {
	resp := &pb.GetWxGzhAccountInfoReply{}
	data, err := a.wxGzhAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = wxGzhAccountInfoFromModel(data)
	return resp, nil
}
