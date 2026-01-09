package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetAiAudioRecordList AI 音乐表-列表数据查询
func (a *AdminV1AiAudioRecordService) GetAiAudioRecordList(ctx context.Context, req *pb.GetAiAudioRecordListReq) (*pb.GetAiAudioRecordListReply, error) {
	resp := &pb.GetAiAudioRecordListReply{
		Total: 0,
		List:  []*pb.AiAudioRecordInfo{},
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
	list, p, err := a.aiAudioRecordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiAudioRecordInfo{
				Id:           v.ID,
				TenantId:     v.TenantID,
				AdminId:      v.AdminID,
				Title:        v.Title,
				Lyric:        v.Lyric,
				ImageURL:     v.ImageURL,
				AudioURL:     v.AudioURL,
				Status:       v.Status,
				Description:  v.Description,
				Prompt:       v.Prompt,
				Platform:     v.Platform,
				ModelId:      v.ModelID,
				Model:        v.Model,
				GenerateMode: v.GenerateMode,
				Tags:         v.Tags,
				Duration:     v.Duration,
				PublicStatus: v.PublicStatus,
				TaskId:       v.TaskID,
				ErrorMessage: v.ErrorMessage,
				CreatedAt:    v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
