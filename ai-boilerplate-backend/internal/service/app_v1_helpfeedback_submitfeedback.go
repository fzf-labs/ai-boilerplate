package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/uuidutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// SubmitFeedback 用户反馈表-提交反馈
func (a *AppV1HelpFeedbackService) SubmitFeedback(ctx context.Context, req *pb.SubmitFeedbackReq) (*pb.SubmitFeedbackReply, error) {
	resp := &pb.SubmitFeedbackReply{}

	// 从上下文获取用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)

	// 创建反馈记录
	feedback := &ai_boilerplate_model.HelpFeedback{
		ID:          uuidutil.GenUUID(),
		UserID:      userID,
		Category:    req.Category,
		Description: req.Description,
		Images:      normalizeHelpFeedbackImages(req.GetImages()),
		Contact:     req.Contact,
		Status:      0, // 默认为待处理
	}

	// 保存到数据库
	err := a.helpFeedbackRepo.CreateOne(ctx, feedback)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Id = feedback.ID
	return resp, nil
}
