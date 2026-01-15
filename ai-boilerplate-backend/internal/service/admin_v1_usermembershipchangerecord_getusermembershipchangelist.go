package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetUserMembershipChangeList 用户会员变更记录表-列表数据查询
func (a *AdminV1UserMembershipChangeRecordService) GetUserMembershipChangeList(ctx context.Context, req *pb.GetUserMembershipChangeListReq) (*pb.GetUserMembershipChangeListReply, error) {
	resp := &pb.GetUserMembershipChangeListReply{
		Total: 0,
		List:  []*pb.UserMembershipChangeInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "user_id",
				Value: req.GetUserId(),
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
	list, p, err := a.userMembershipChangeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	for _, v := range list {
		before := &ai_boilerplate_model.UserMembership{}
		if v.Before.String() != "" {
			err = jsonutil.Unmarshal(v.Before, before)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		after := &ai_boilerplate_model.UserMembership{}
		if v.After.String() != "" {
			err = jsonutil.Unmarshal(v.After, after)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		resp.List = append(resp.List, &pb.UserMembershipChangeInfo{
			Id:         v.ID,
			UserId:     v.UserID,
			SourceType: v.SourceType,
			SourceId:   v.SourceID,
			Before: &pb.UserMembershipChangeItem{
				UserId:         before.UserID,
				MembershipType: before.MembershipType,
				ExpiredAt:      timeutil.RFC3339(before.ExpiredAt.Time),
				AutoRenew:      before.AutoRenew,
				AutoRenewDays:  before.AutoRenewDays,
				Status:         before.Status,
				CreatedAt:      before.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      before.UpdatedAt.Format(time.RFC3339),
			},
			After: &pb.UserMembershipChangeItem{
				UserId:         after.UserID,
				MembershipType: after.MembershipType,
				ExpiredAt:      timeutil.RFC3339(after.ExpiredAt.Time),
				AutoRenew:      after.AutoRenew,
				AutoRenewDays:  after.AutoRenewDays,
				Status:         after.Status,
				CreatedAt:      after.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      after.UpdatedAt.Format(time.RFC3339),
			},
			Remark:    v.Remark,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
		})
	}
	return resp, nil
}
