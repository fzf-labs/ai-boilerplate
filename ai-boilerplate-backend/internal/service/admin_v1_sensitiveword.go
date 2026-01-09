package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1SensitiveWordService(
	logger log.Logger,
	sensitiveWordRepo *data.SensitiveWordRepo,
) *AdminV1SensitiveWordService {
	l := log.NewHelper(log.With(logger, "module", "service/sensitiveWord"))
	return &AdminV1SensitiveWordService{
		log:               l,
		sensitiveWordRepo: sensitiveWordRepo,
	}
}

type AdminV1SensitiveWordService struct {
	pb.UnimplementedSensitiveWordServer
	log               *log.Helper
	sensitiveWordRepo *data.SensitiveWordRepo
}
