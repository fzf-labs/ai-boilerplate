package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"gorm.io/datatypes"
)

// CreateSysOperateLog 系统-操作日志-创建一条数据
func (a *AdminV1SysOperateLogService) CreateSysOperateLog(ctx context.Context, req *pb.CreateSysOperateLogReq) (*pb.CreateSysOperateLogReply, error) {
	resp := &pb.CreateSysOperateLogReply{}
	data := a.sysOperateLogRepo.NewData()
	data.TenantID = req.GetTenantId()
	data.TraceID = req.GetTraceId()
	data.AdminID = req.GetAdminId()
	data.IP = req.GetIP()
	data.URI = req.GetURI()
	data.Useragent = req.GetUseragent()
	data.Header = datatypes.JSON(req.GetHeader())
	data.Req = datatypes.JSON(req.Req)
	data.Resp = datatypes.JSON(req.Resp)
	err := a.sysOperateLogRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
