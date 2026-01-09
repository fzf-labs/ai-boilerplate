package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/samber/lo"
)

// GetSysTenantList 系统-租户-列表数据查询
func (a *AdminV1SysTenantService) GetSysTenantList(ctx context.Context, req *pb.GetSysTenantListReq) (*pb.GetSysTenantListReply, error) {
	resp := &pb.GetSysTenantListReply{
		Total: 0,
		List:  []*pb.SysTenantInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.ASC,
			},
		},
	}
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
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
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query,
			&condition.QueryParam{
				Field: "created_at",
				Value: req.GetCreatedAt()[0],
				Exp:   condition.GTE,
				Logic: condition.AND,
			}, &condition.QueryParam{
				Field: "created_at",
				Value: req.GetCreatedAt()[1],
				Exp:   condition.LTE,
				Logic: condition.AND,
			})
	}
	list, p, err := a.sysTenantRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		adminIDs := lo.Map(list, func(item *ai_boilerplate_model.SysTenant, _ int) string {
			return item.AdminID
		})
		adminNameMap, err := a.sysAdminRepo.AdminIDToNickname(ctx, adminIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}

		for _, v := range list {
			menuIDs := make([]string, 0)
			if v.MenuIds.String() != "" {
				err = jsonutil.Unmarshal(v.MenuIds, &menuIDs)
				if err != nil {
					return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
				}
			}
			resp.List = append(resp.List, &pb.SysTenantInfo{
				Id:         v.ID,
				Name:       v.Name,
				Remark:     v.Remark,
				AdminId:    v.AdminID,
				ExpireTime: timeutil.RFC3339(v.ExpireTime.Time),
				MenuIds:    menuIDs,
				Status:     int32(v.Status),
				CreatedAt:  v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:  v.UpdatedAt.Format(time.RFC3339),
				AdminName:  adminNameMap[v.AdminID],
			})
		}
	}
	return resp, nil
}
