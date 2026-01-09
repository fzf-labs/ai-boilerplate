package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateSysPost 系统-工作岗位-创建一条数据
func (a *AdminV1SysPostService) CreateSysPost(ctx context.Context, req *pb.CreateSysPostReq) (*pb.CreateSysPostReply, error) {
	resp := &pb.CreateSysPostReply{}
	data := a.sysPostRepo.NewData()
	data.TenantID = meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data.Name = req.GetName()
	data.Code = req.GetCode()
	data.Remark = req.GetRemark()
	data.Sort = int64(req.GetSort())
	data.Status = int16(req.GetStatus())
	err := a.sysPostRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
