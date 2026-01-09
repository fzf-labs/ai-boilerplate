package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSysAPIList 系统-接口-列表数据查询
func (a *AdminV1SysAPIService) GetSysAPIList(ctx context.Context, req *pb.GetSysAPIListReq) (*pb.GetSysAPIListReply, error) {
	resp := &pb.GetSysAPIListReply{
		Total: 0,
		List:  []*pb.SysAPIInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.sysAPIRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysAPIInfo{
				Id: v.ID,
				// TODO
			})
		}
	}
	return resp, nil
}
