package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateAiWriteRecord AI 写作表-创建一条数据
func (a *AdminV1AiWriteRecordService) CreateAiWriteRecord(ctx context.Context, req *pb.CreateAiWriteRecordReq) (*pb.CreateAiWriteRecordReply, error) {
	resp := &pb.CreateAiWriteRecordReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data := a.aiWriteRecordRepo.NewData()
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
	err := a.aiWriteRecordRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
