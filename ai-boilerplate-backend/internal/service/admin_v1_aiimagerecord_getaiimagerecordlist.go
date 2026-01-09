package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetAiImageRecordList AI 绘画表-列表数据查询
func (a *AdminV1AiImageRecordService) GetAiImageRecordList(ctx context.Context, req *pb.GetAiImageRecordListReq) (*pb.GetAiImageRecordListReply, error) {
	resp := &pb.GetAiImageRecordListReply{
		Total: 0,
		List:  []*pb.AiImageRecordInfo{},
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
	list, p, err := a.aiImageRecordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiImageRecordInfo{
				Id:           v.ID,
				AdminId:      v.AdminID,
				Prompt:       v.Prompt,
				Platform:     v.Platform,
				ModelId:      v.ModelID,
				Model:        v.Model,
				Width:        v.Width,
				Height:       v.Height,
				Status:       v.Status,
				FinishTime:   v.FinishTime.Time.Format(time.RFC3339),
				ErrorMessage: v.ErrorMessage,
				PublicStatus: v.PublicStatus,
				PicURL:       v.PicURL,
				Options:      v.Options.String(),
				TaskId:       v.TaskID,
				Buttons:      v.Buttons,
				CreatedAt:    v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
