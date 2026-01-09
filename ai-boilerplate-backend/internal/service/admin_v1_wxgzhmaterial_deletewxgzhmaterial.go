package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxGzhMaterial 公众号素材表-删除一条数据
func (a *AdminV1WxGzhMaterialService) DeleteWxGzhMaterial(ctx context.Context, req *pb.DeleteWxGzhMaterialReq) (*pb.DeleteWxGzhMaterialReply, error) {
	resp := &pb.DeleteWxGzhMaterialReply{}
	err := a.wxGzhMaterialRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
