package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysAdminList 系统-用户-列表数据查询
func (a *AdminV1SysAdminService) GetSysAdminList(ctx context.Context, req *pb.GetSysAdminListReq) (*pb.GetSysAdminListReply, error) {
	resp := &pb.GetSysAdminListReply{
		Total: 0,
		List:  []*pb.SysAdminInfo{},
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
	if req.GetUsername() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "username",
			Value: "%" + req.GetUsername() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetNickname() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "nickname",
			Value: "%" + req.GetNickname() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetMobile() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "mobile",
			Value: req.GetMobile(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetDeptId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "dept_id",
			Value: req.GetDeptId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) >= 2 && req.GetCreatedAt()[0] != "" && req.GetCreatedAt()[1] != "" {
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
	list, p, err := a.sysAdminRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		roleIDs := make([]string, 0)
		deptIDs := make([]string, 0)
		postIDs := make([]string, 0)
		for _, v := range list {
			roleIDs = append(roleIDs, v.RoleID)
			deptIDs = append(deptIDs, v.DeptID)
			postIDs = append(postIDs, v.PostID)
		}
		roleNameMap, err := a.sysRoleRepo.RoleIDToName(ctx, roleIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		deptNameMap, err := a.sysDeptRepo.DeptIDToName(ctx, deptIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		postNameMap, err := a.sysPostRepo.PostIDToName(ctx, postIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysAdminInfo{
				Id:        v.ID,
				Username:  v.Username,
				Nickname:  v.Nickname,
				Avatar:    v.Avatar,
				Sex:       int32(v.Sex),
				Email:     v.Email,
				Mobile:    v.Mobile,
				RoleId:    v.RoleID,
				DeptId:    v.DeptID,
				PostId:    v.PostID,
				Status:    int32(v.Status),
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
				RoleName:  roleNameMap[v.RoleID],
				DeptName:  deptNameMap[v.DeptID],
				PostName:  postNameMap[v.PostID],
			})
		}
	}
	return resp, nil
}
