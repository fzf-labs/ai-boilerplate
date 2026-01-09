package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysNoticeList 系统-公告-列表数据查询
func (a *AdminV1SysNoticeService) GetSysNoticeList(ctx context.Context, req *pb.GetSysNoticeListReq) (*pb.GetSysNoticeListReply, error) {
	resp := &pb.GetSysNoticeListReply{
		Total: 0,
		List:  []*pb.SysNoticeInfo{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "tenant_id",
				Value: tenantID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	if req.GetType() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "type",
			Value: req.GetType(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetTitle() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "title",
			Value: "%" + req.GetTitle() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.sysNoticeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysNoticeInfo{
				Id:        v.ID,
				Type:      v.Type,
				Title:     v.Title,
				Content:   v.Content,
				Status:    int32(v.Status),
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
