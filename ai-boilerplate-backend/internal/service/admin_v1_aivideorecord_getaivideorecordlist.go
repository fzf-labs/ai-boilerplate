package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetAiVideoRecordList AI 视频表-列表数据查询
func (a *AdminV1AiVideoRecordService) GetAiVideoRecordList(ctx context.Context, req *pb.GetAiVideoRecordListReq) (*pb.GetAiVideoRecordListReply, error) {
	resp := &pb.GetAiVideoRecordListReply{
		Total: 0,
		List:  []*pb.AiVideoRecordInfo{},
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
	list, p, err := a.aiVideoRecordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiVideoRecordInfo{
				Id:           v.ID,
				AdminId:      v.AdminID,
				Prompt:       v.Prompt,
				Platform:     v.Platform,
				ModelId:      v.ModelID,
				Model:        v.Model,
				Status:       v.Status,
				FinishTime:   v.FinishTime.Time.Format(time.RFC3339),
				ErrorMessage: v.ErrorMessage,
				PublicStatus: v.PublicStatus,
				VideoURL:     v.VideoURL,
				Options:      v.Options.String(),
				TaskId:       v.TaskID,
				CreatedAt:    v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
