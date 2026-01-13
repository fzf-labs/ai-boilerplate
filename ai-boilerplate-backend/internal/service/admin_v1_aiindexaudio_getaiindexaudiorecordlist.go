package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexAudioRecordList AI 音乐表-列表数据查询
func (a *AdminV1AiIndexAudioService) GetAiIndexAudioRecordList(ctx context.Context, req *pb.GetAiIndexAudioRecordListReq) (*pb.GetAiIndexAudioRecordListReply, error) {
	resp := &pb.GetAiIndexAudioRecordListReply{
		Total: 0,
		List:  []*pb.AiIndexAudioRecordInfo{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
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
	if tenantID != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "tenant_id",
			Value: tenantID,
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if adminID != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "admin_id",
			Value: adminID,
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.aiAudioRecordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiIndexAudioRecordInfo{
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
