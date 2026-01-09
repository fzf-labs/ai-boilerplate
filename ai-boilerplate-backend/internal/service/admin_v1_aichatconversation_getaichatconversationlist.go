package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetAiChatConversationList AI 聊天对话表-列表数据查询
func (a *AdminV1AiChatConversationService) GetAiChatConversationList(ctx context.Context, req *pb.GetAiChatConversationListReq) (*pb.GetAiChatConversationListReply, error) {
	resp := &pb.GetAiChatConversationListReply{
		Total: 0,
		List:  []*pb.AiChatConversationInfo{},
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
	list, p, err := a.aiChatConversationRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			promptSetting := &pb.AiChatConversationInfo_PromptSetting{}
			err = jsonutil.Unmarshal(v.PromptSetting, promptSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			modelSetting := &pb.AiChatConversationInfo_ModelSetting{}
			err = jsonutil.Unmarshal(v.ModelSetting, modelSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			knowledgeSetting := &pb.AiChatConversationInfo_KnowledgeSetting{}
			err = jsonutil.Unmarshal(v.KnowledgeSetting, knowledgeSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			mcpSetting := &pb.AiChatConversationInfo_McpSetting{}
			err = jsonutil.Unmarshal(v.McpSetting, mcpSetting)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			resp.List = append(resp.List, &pb.AiChatConversationInfo{
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
