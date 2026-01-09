package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/samber/lo"
)

// GetSysDeptList 系统-部门-列表数据查询
func (a *AdminV1SysDeptService) GetSysDeptList(ctx context.Context, _ *pb.GetSysDeptListReq) (*pb.GetSysDeptListReply, error) {
	resp := &pb.GetSysDeptListReply{
		List: []*pb.SysDeptInfo{},
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
	list, _, err := a.sysDeptRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) > 0 {
		adminIDs := lo.Map(list, func(item *ai_boilerplate_model.SysDept, _ int) string {
			return item.AdminID
		})
		adminIDToNickname, err := a.sysAdminRepo.AdminIDToNickname(ctx, adminIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysDeptInfo{
				Id:        v.ID,
				Pid:       v.Pid,
				Name:      v.Name,
				AdminId:   v.AdminID,
				Status:    int32(v.Status),
				Sort:      int32(v.Sort),
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
				AdminName: adminIDToNickname[v.AdminID],
			})
		}
	}
	resp.Total = int32(len(resp.List))
	return resp, nil
}
