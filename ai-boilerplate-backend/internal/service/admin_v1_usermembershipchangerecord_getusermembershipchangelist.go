package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
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
		before := &pb.UserMembershipChangeItem{}
		if v.Before.String() != "" {
			err = jsonutil.Unmarshal(v.Before, before)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		after := &pb.UserMembershipChangeItem{}
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
			Before:     before,
			After:      after,
			Remark:     v.Remark,
			CreatedAt:  v.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  v.UpdatedAt.Format(time.RFC3339),
		})
	}
	return resp, nil
}
