package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetWxXcxUserList 小程序用户表-列表数据查询
func (a *AdminV1WxXcxUserService) GetWxXcxUserList(ctx context.Context, req *pb.GetWxXcxUserListReq) (*pb.GetWxXcxUserListReply, error) {
	resp := &pb.GetWxXcxUserListReply{
		Total: 0,
		List:  []*pb.WxXcxUserInfo{},
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
	list, p, err := a.wxXcxUserRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.WxXcxUserInfo{
				Id:        v.ID,
				AppId:     v.AppID,
				Openid:    v.Openid,
				Unionid:   v.Unionid,
				Nickname:  v.Nickname,
				AvatarURL: v.AvatarURL,
				Language:  v.Language,
				Country:   v.Country,
				Province:  v.Province,
				City:      v.City,
				Remark:    v.Remark,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
