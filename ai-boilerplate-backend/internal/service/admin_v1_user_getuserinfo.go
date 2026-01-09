package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetUserInfo 用户表-单条数据查询
func (a *AdminV1UserService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoReply, error) {
	resp := &pb.GetUserInfoReply{
		Info: &pb.UserInfo{},
	}
	data, err := a.userRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return resp, nil
	}
	userMembershipInfo := &pb.UserMembershipInfo{}
	userMembership, err := a.userMembershipRepo.FindOneByUserID(ctx, data.ID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if userMembership != nil && userMembership.ID != "" {
		userMembershipInfo = &pb.UserMembershipInfo{
			Id:             userMembership.ID,
			UserId:         userMembership.UserID,
			MembershipType: userMembership.MembershipType,
			ExpiredAt:      userMembership.ExpiredAt.Time.Format(time.RFC3339),
			AutoRenew:      userMembership.AutoRenew,
			AutoRenewDays:  userMembership.AutoRenewDays,
			Status:         userMembership.Status,
			CreatedAt:      userMembership.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      userMembership.UpdatedAt.Format(time.RFC3339),
		}
	}
	resp.Info = &pb.UserInfo{
		Id:                 data.ID,
		Phone:              data.Phone,
		Nickname:           data.Nickname,
		Gender:             data.Gender,
		Avatar:             data.Avatar,
		Profile:            data.Profile,
		WxGzhUserId:        data.WxGzhUserID,
		WxGzhXcxId:         data.WxGzhXcxID,
		Status:             data.Status,
		CreatedAt:          data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          data.UpdatedAt.Format(time.RFC3339),
		UserMembershipInfo: userMembershipInfo,
	}
	return resp, nil
}
