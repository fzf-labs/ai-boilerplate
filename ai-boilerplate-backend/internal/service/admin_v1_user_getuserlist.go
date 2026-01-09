package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetUserList 用户表-列表数据查询
func (a *AdminV1UserService) GetUserList(ctx context.Context, req *pb.GetUserListReq) (*pb.GetUserListReply, error) {
	resp := &pb.GetUserListReply{
		Total: 0,
		List:  []*pb.UserInfo{},
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
	if req.GetPhone() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "phone",
			Value: "%" + req.GetPhone() + "%",
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
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
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
	list, p, err := a.userRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		userIDs := make([]string, 0)
		for _, v := range list {
			userIDs = append(userIDs, v.ID)
		}
		userMembershipList, err := a.userMembershipRepo.FindMultiCacheByUserIDS(ctx, userIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		userMembershipMap := make(map[string]*pb.UserMembershipInfo)
		for _, v := range userMembershipList {
			userMembershipMap[v.UserID] = &pb.UserMembershipInfo{
				Id:             v.ID,
				UserId:         v.UserID,
				MembershipType: v.MembershipType,
				ExpiredAt:      v.ExpiredAt.Time.Format(time.RFC3339),
				AutoRenew:      v.AutoRenew,
				AutoRenewDays:  v.AutoRenewDays,
				Status:         v.Status,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
			}
		}
		for _, v := range list {
			resp.List = append(resp.List, &pb.UserInfo{
				Id:                 v.ID,
				Phone:              v.Phone,
				Nickname:           v.Nickname,
				Gender:             v.Gender,
				Avatar:             v.Avatar,
				Profile:            v.Profile,
				WxGzhUserId:        v.WxGzhUserID,
				WxGzhXcxId:         v.WxGzhXcxID,
				Status:             v.Status,
				CreatedAt:          v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:          v.UpdatedAt.Format(time.RFC3339),
				UserMembershipInfo: userMembershipMap[v.ID],
			})
		}
	}
	return resp, nil
}
