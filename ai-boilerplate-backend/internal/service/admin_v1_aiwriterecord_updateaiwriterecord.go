package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateAiWriteRecord AI 写作表-更新一条数据
func (a *AdminV1AiWriteRecordService) UpdateAiWriteRecord(ctx context.Context, req *pb.UpdateAiWriteRecordReq) (*pb.UpdateAiWriteRecordReply, error) {
	resp := &pb.UpdateAiWriteRecordReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data, err := a.aiWriteRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiWriteRecordRepo.DeepCopy(data)
	data.TenantID = tenantID
	data.AdminID = req.GetAdminId()
	data.Type = req.GetType()
	data.Platform = req.GetPlatform()
	data.ModelID = req.GetModelId()
	data.Model = req.GetModel()
	data.Prompt = req.GetPrompt()
	data.GeneratedContent = req.GetGeneratedContent()
	data.OriginalContent = req.GetOriginalContent()
	data.Length = req.GetLength()
	data.Format = req.GetFormat()
	data.Tone = req.GetTone()
	data.Language = req.GetLanguage()
	data.ErrorMessage = req.GetErrorMessage()
	err = a.aiWriteRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
