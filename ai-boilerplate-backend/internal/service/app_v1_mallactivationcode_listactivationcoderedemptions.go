package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// ListActivationCodeRedemptions 会员激活码兑换记录
func (a *AppV1MallActivationCodeService) ListActivationCodeRedemptions(ctx context.Context, req *pb.ListActivationCodeRedemptionsReq) (*pb.ListActivationCodeRedemptionsReply, error) {
	resp := &pb.ListActivationCodeRedemptionsReply{
		Total: 0,
		List:  []*pb.ActivationCodeRedemptionInfo{},
	}

	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	if userID == "" {
		return nil, pb.ErrorReasonUnauthorized()
	}

	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "user_id",
				Value: userID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "status",
				Value: int32(constant.ActivationCodeStatusActivated),
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{Field: "activated_at", Order: condition.DESC},
		},
	}

	list, p, err := a.mallActivationCodeRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total

	for _, item := range list {
		info := &pb.ActivationCodeRedemptionInfo{
			Code: item.Code,
		}
		if item.ActivatedAt.Valid {
			info.ActivatedAt = item.ActivatedAt.Time.Format(time.RFC3339)
		}

		userChange := &activationCodeUserChange{}
		if item.UserChange.String() != "" {
			if err := jsonutil.Unmarshal(item.UserChange, userChange); err == nil {
				if userChange.UserMembershipChange != nil && userChange.UserMembershipChange.After != nil {
					info.MembershipType = userChange.UserMembershipChange.After.MembershipType
					info.ExpiredAt = userChange.UserMembershipChange.After.ExpiredAt
					info.DurationDays = userChange.UserMembershipChange.After.DurationDays
				}
			}
		}
		resp.List = append(resp.List, info)
	}

	return resp, nil
}
