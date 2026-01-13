package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexImageRecordList AI 绘画表-列表数据查询
func (a *AdminV1AiIndexImageService) GetAiIndexImageRecordList(ctx context.Context, req *pb.GetAiIndexImageRecordListReq) (*pb.GetAiIndexImageRecordListReply, error) {
	resp := &pb.GetAiIndexImageRecordListReply{
		Total: 0,
		List:  []*pb.AiIndexImageRecordInfo{},
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
	list, p, err := a.aiImageRecordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			finishTime := ""
			if v.FinishTime.Valid {
				finishTime = v.FinishTime.Time.Format(time.RFC3339)
			}
			resp.List = append(resp.List, &pb.AiIndexImageRecordInfo{
				Id:           v.ID,
				AdminId:      v.AdminID,
				Prompt:       v.Prompt,
				Platform:     v.Platform,
				ModelId:      v.ModelID,
				Model:        v.Model,
				Width:        v.Width,
				Height:       v.Height,
				Status:       v.Status,
				FinishTime:   finishTime,
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
