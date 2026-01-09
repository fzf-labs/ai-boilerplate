package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserInfo 用户表-单条数据查询
func (a *AppV1UserService) GetUserInfo(ctx context.Context, _ *pb.GetUserInfoReq) (*pb.GetUserInfoReply, error) {
	resp := &pb.GetUserInfoReply{}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	data, err := a.userRepo.FindOneCacheByID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.UserInfo{
		Id:          data.ID,
		Phone:       data.Phone,
		Nickname:    data.Nickname,
		Gender:      data.Gender,
		Avatar:      data.Avatar,
		Profile:     data.Profile,
		WxGzhUserId: data.WxGzhUserID,
		WxGzhXcxId:  data.WxGzhXcxID,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
