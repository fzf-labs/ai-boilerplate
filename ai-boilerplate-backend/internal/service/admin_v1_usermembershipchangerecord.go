package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1UserMembershipChangeRecordService(
	logger log.Logger,
	userMembershipChangeRepo *data.UserMembershipChangeRepo,
) *AdminV1UserMembershipChangeRecordService {
	l := log.NewHelper(log.With(logger, "module", "service/userMembershipChangeRecord"))
	return &AdminV1UserMembershipChangeRecordService{
		log:                      l,
		userMembershipChangeRepo: userMembershipChangeRepo,
	}
}

type AdminV1UserMembershipChangeRecordService struct {
	pb.UnimplementedUserMembershipChangeRecordServer
	log                      *log.Helper
	userMembershipChangeRepo *data.UserMembershipChangeRepo
}
