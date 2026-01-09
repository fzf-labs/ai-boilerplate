package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSysPostSelector 系统-工作岗位-选择器
func (a *AdminV1SysPostService) GetSysPostSelector(ctx context.Context, _ *pb.GetSysPostSelectorReq) (*pb.GetSysPostSelectorReply, error) {
	resp := &pb.GetSysPostSelectorReply{
		List: []*pb.SysPostSelectorItem{},
	}
	param := &condition.Req{
		Query: []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "sort",
				Order: condition.ASC,
			},
		},
	}
	list, _, err := a.sysPostRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysPostSelectorItem{
				Id:   v.ID,
				Name: v.Name,
			})
		}
	}
	return resp, nil
}
