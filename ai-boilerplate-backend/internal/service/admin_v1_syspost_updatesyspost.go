package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateSysPost 系统-工作岗位-更新一条数据
func (a *AdminV1SysPostService) UpdateSysPost(ctx context.Context, req *pb.UpdateSysPostReq) (*pb.UpdateSysPostReply, error) {
	resp := &pb.UpdateSysPostReply{}
	data, err := a.sysPostRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysPostRepo.DeepCopy(data)
	data.TenantID = meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data.Name = req.GetName()
	data.Code = req.GetCode()
	data.Remark = req.GetRemark()
	data.Sort = int64(req.GetSort())
	data.Status = int16(req.GetStatus())
	err = a.sysPostRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
