package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexChatConversationList AI 聊天对话表-列表数据查询
func (a *AdminV1AiIndexChatService) GetAiIndexChatConversationList(ctx context.Context, req *pb.GetAiIndexChatConversationListReq) (*pb.GetAiIndexChatConversationListReply, error) {
	resp := &pb.GetAiIndexChatConversationListReply{
		Total: 0,
		List:  []*pb.AiIndexChatConversationItem{},
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
	list, p, err := a.aiChatConversationRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			promptSetting := &pb.AiIndexChatConversationItem_PromptSetting{}
			err = jsonutil.Unmarshal(v.PromptSetting, promptSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			modelSetting := &pb.AiIndexChatConversationItem_ModelSetting{}
			err = jsonutil.Unmarshal(v.ModelSetting, modelSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			knowledgeSetting := &pb.AiIndexChatConversationItem_KnowledgeSetting{}
			err = jsonutil.Unmarshal(v.KnowledgeSetting, knowledgeSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			mcpSetting := &pb.AiIndexChatConversationItem_McpSetting{}
			err = jsonutil.Unmarshal(v.McpSetting, mcpSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			resp.List = append(resp.List, &pb.AiIndexChatConversationItem{
				Id:               v.ID,
				AdminId:          v.AdminID,
				Title:            v.Title,
				Pinned:           v.Pinned,
				PinnedTime:       v.PinnedTime.Time.Format(time.RFC3339),
				PromptSetting:    promptSetting,
				ModelSetting:     modelSetting,
				KnowledgeSetting: knowledgeSetting,
				McpSetting:       mcpSetting,
				CreatedAt:        v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:        v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
