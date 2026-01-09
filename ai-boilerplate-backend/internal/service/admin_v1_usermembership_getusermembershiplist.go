package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetUserMembershipList 用户会员关系表-列表数据查询
func (a *AdminV1UserMembershipService) GetUserMembershipList(ctx context.Context, req *pb.GetUserMembershipListReq) (*pb.GetUserMembershipListReply, error) {
	resp := &pb.GetUserMembershipListReply{
		Total: 0,
		List:  []*pb.UserMembershipInfo{},
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
	list, p, err := a.userMembershipRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.UserMembershipInfo{
				Id:             v.ID,
				UserId:         v.UserID,
				MembershipType: v.MembershipType,
				ExpiredAt:      v.ExpiredAt.Time.Format(time.RFC3339),
				AutoRenew:      v.AutoRenew,
				AutoRenewDays:  v.AutoRenewDays,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
