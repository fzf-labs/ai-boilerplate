package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexChatMessageList AI 聊天消息表-列表数据查询
func (a *AdminV1AiIndexChatService) GetAiIndexChatMessageList(ctx context.Context, req *pb.GetAiIndexChatMessageListReq) (*pb.GetAiIndexChatMessageListReply, error) {
	resp := &pb.GetAiIndexChatMessageListReply{
		Total: 0,
		List:  []*pb.AiIndexChatMessageItem{},
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
	list, p, err := a.aiChatMessageRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiIndexChatMessageItem{
				Id:             v.ID,
				ConversationId: v.ConversationID,
				ReplyId:        v.ReplyID,
				AdminId:        v.AdminID,
				RoleId:         v.RoleID,
				Type:           v.Type,
				Model:          v.Model,
				ModelId:        v.ModelID,
				Content:        v.Content,
				UseContext:     v.UseContext,
				SegmentIds:     v.SegmentIds,
				CreatedAt:      v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
