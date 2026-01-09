package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1AiWriteRecordService(
	logger log.Logger,
	aiWriteRecordRepo *data.AiWriteRecordRepo,
) *AdminV1AiWriteRecordService {
	l := log.NewHelper(log.With(logger, "module", "service/aiWriteRecord"))
	return &AdminV1AiWriteRecordService{
		log:               l,
		aiWriteRecordRepo: aiWriteRecordRepo,
	}
}

type AdminV1AiWriteRecordService struct {
	pb.UnimplementedAiWriteRecordServer
	log               *log.Helper
	aiWriteRecordRepo *data.AiWriteRecordRepo
}
