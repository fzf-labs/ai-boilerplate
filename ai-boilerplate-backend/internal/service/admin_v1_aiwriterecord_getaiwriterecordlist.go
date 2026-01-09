package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetAiWriteRecordList AI 写作表-列表数据查询
func (a *AdminV1AiWriteRecordService) GetAiWriteRecordList(ctx context.Context, req *pb.GetAiWriteRecordListReq) (*pb.GetAiWriteRecordListReply, error) {
	resp := &pb.GetAiWriteRecordListReply{
		Total: 0,
		List:  []*pb.AiWriteRecordInfo{},
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
	list, p, err := a.aiWriteRecordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiWriteRecordInfo{
				Id:               v.ID,
				AdminId:          v.AdminID,
				Type:             v.Type,
				Platform:         v.Platform,
				ModelId:          v.ModelID,
				Model:            v.Model,
				Prompt:           v.Prompt,
				GeneratedContent: v.GeneratedContent,
				OriginalContent:  v.OriginalContent,
				Length:           v.Length,
				Format:           v.Format,
				Tone:             v.Tone,
				Language:         v.Language,
				ErrorMessage:     v.ErrorMessage,
				CreatedAt:        v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:        v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
