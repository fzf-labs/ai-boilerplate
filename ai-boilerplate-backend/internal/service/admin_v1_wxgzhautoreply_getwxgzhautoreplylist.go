package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetWxGzhAutoReplyList 公众号消息自动回复表-列表数据查询
func (a *AdminV1WxGzhAutoReplyService) GetWxGzhAutoReplyList(ctx context.Context, req *pb.GetWxGzhAutoReplyListReq) (*pb.GetWxGzhAutoReplyListReply, error) {
	resp := &pb.GetWxGzhAutoReplyListReply{
		Total: 0,
		List:  []*pb.WxGzhAutoReplyInfo{},
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
	if req.GetAppId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "app_id",
			Value: req.GetAppId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetType() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "type",
			Value: req.GetType(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.wxGzhAutoReplyRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.WxGzhAutoReplyInfo{
				Id:                  v.ID,
				AppId:               v.AppID,
				Type:                v.Type,
				RequestKeyword:      v.RequestKeyword,
				RequestKeywordMatch: v.RequestKeywordMatch,
				ResponseMessageType: v.ResponseMessageType,
				ResponseContent:     v.ResponseContent,
				ResponseMediaId:     v.ResponseMediaID,
				Status:              v.Status,
				CreatedAt:           v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:           v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
