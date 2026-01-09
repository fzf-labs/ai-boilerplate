package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysOperateLogList 系统-操作日志-列表数据查询
func (a *AdminV1SysOperateLogService) GetSysOperateLogList(ctx context.Context, req *pb.GetSysOperateLogListReq) (*pb.GetSysOperateLogListReply, error) {
	resp := &pb.GetSysOperateLogListReply{
		Total: 0,
		List:  []*pb.SysOperateLogInfo{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "tenant_id",
				Value: tenantID,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetTraceId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "trace_id",
			Value: req.GetTraceId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetAdminId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "admin_id",
			Value: req.GetAdminId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[0],
			Exp:   condition.GTE,
			Logic: condition.AND,
		})
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[1],
			Exp:   condition.LTE,
			Logic: condition.AND,
		})
	}
	list, p, err := a.sysOperateLogRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		adminIDs := make([]string, 0)
		for _, v := range list {
			adminIDs = append(adminIDs, v.AdminID)
		}
		adminMap, err := a.sysAdminRepo.AdminIDToNickname(ctx, adminIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysOperateLogInfo{
				Id:        v.ID,
				AdminId:   v.AdminID,
				IP:        v.IP,
				URI:       v.URI,
				Useragent: v.Useragent,
				Header:    v.Header.String(),
				Req:       v.Req.String(),
				Resp:      v.Resp.String(),
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				TraceId:   v.TraceID,
				Nickname:  adminMap[v.AdminID],
			})
		}
	}
	return resp, nil
}
