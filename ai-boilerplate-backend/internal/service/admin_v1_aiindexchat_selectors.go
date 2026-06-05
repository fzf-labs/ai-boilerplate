package service

import (
	"context"
	"sort"
	"strings"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexChatKnowledgeSelector AI 聊天对话表-知识库选择器
func (a *AdminV1AiIndexChatService) GetAiIndexChatKnowledgeSelector(ctx context.Context, _ *pb.GetAiIndexChatKnowledgeSelectorReq) (*pb.GetAiIndexChatKnowledgeSelectorReply, error) {
	conversations, err := a.getAiIndexChatConversationSelectorList(ctx)
	if err != nil {
		return nil, err
	}
	knowledgeIDs, err := extractAiIndexChatKnowledgeIDs(conversations)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetAiIndexChatKnowledgeSelectorReply{
		List: make([]*pb.AiIndexChatKnowledgeSelectorItem, 0, len(knowledgeIDs)),
	}
	for _, id := range knowledgeIDs {
		resp.List = append(resp.List, &pb.AiIndexChatKnowledgeSelectorItem{
			Label: id,
			Value: id,
		})
	}
	return resp, nil
}

// GetAiIndexChatMcpSelector AI 聊天对话表-MCP 选择器
func (a *AdminV1AiIndexChatService) GetAiIndexChatMcpSelector(ctx context.Context, _ *pb.GetAiIndexChatMcpSelectorReq) (*pb.GetAiIndexChatMcpSelectorReply, error) {
	conversations, err := a.getAiIndexChatConversationSelectorList(ctx)
	if err != nil {
		return nil, err
	}
	mcpIDs, err := extractAiIndexChatMcpIDs(conversations)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetAiIndexChatMcpSelectorReply{
		List: make([]*pb.AiIndexChatMcpSelectorItem, 0, len(mcpIDs)),
	}
	for _, id := range mcpIDs {
		resp.List = append(resp.List, &pb.AiIndexChatMcpSelectorItem{
			Label: id,
			Value: id,
		})
	}
	return resp, nil
}

func (a *AdminV1AiIndexChatService) getAiIndexChatConversationSelectorList(ctx context.Context) ([]*pb.AiIndexChatConversationItem, error) {
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)

	param := &condition.Req{
		Query: []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "updated_at",
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

	list, _, err := a.aiChatConversationRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp := make([]*pb.AiIndexChatConversationItem, 0, len(list))
	for _, v := range list {
		promptSetting := &pb.AiIndexChatConversationItem_PromptSetting{}
		if len(v.PromptSetting) > 0 {
			if err := jsonutil.Unmarshal(v.PromptSetting, promptSetting); err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		modelSetting := &pb.AiIndexChatConversationItem_ModelSetting{}
		if len(v.ModelSetting) > 0 {
			if err := jsonutil.Unmarshal(v.ModelSetting, modelSetting); err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		knowledgeSetting := &pb.AiIndexChatConversationItem_KnowledgeSetting{}
		if len(v.KnowledgeSetting) > 0 {
			if err := jsonutil.Unmarshal(v.KnowledgeSetting, knowledgeSetting); err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}
		mcpSetting := &pb.AiIndexChatConversationItem_McpSetting{}
		if len(v.McpSetting) > 0 {
			if err := jsonutil.Unmarshal(v.McpSetting, mcpSetting); err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
		}

		resp = append(resp, &pb.AiIndexChatConversationItem{
			Id:               v.ID,
			AdminId:          v.AdminID,
			Title:            v.Title,
			Pinned:           v.Pinned,
			PinnedTime:       v.PinnedTime.Time.Format(time.RFC3339),
			PromptSetting:    promptSetting,
			ModelSetting:     modelSetting,
			KnowledgeSetting: knowledgeSetting,
			McpSetting:       mcpSetting,
			CreatedAt:        v.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:        v.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return resp, nil
}

func extractAiIndexChatKnowledgeIDs(conversations []*pb.AiIndexChatConversationItem) ([]string, error) {
	unique := make(map[string]struct{})
	for _, conversation := range conversations {
		for _, id := range conversation.GetKnowledgeSetting().GetKnowledgeIds() {
			if id = normalizeAiIndexChatSelectorID(id); id != "" {
				unique[id] = struct{}{}
			}
		}
	}
	return sortAiIndexChatSelectorIDs(unique), nil
}

func extractAiIndexChatMcpIDs(conversations []*pb.AiIndexChatConversationItem) ([]string, error) {
	unique := make(map[string]struct{})
	for _, conversation := range conversations {
		for _, id := range conversation.GetMcpSetting().GetMcpIds() {
			if id = normalizeAiIndexChatSelectorID(id); id != "" {
				unique[id] = struct{}{}
			}
		}
	}
	return sortAiIndexChatSelectorIDs(unique), nil
}

func sortAiIndexChatSelectorIDs(unique map[string]struct{}) []string {
	ids := make([]string, 0, len(unique))
	for id := range unique {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func normalizeAiIndexChatSelectorID(value string) string {
	return strings.TrimSpace(value)
}
