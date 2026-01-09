package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAdminV1DeviceService(
	logger log.Logger,
	deviceRepo *data.DeviceRepo,
	deviceHeartbeatRepo *data.DeviceHeartbeatRepo,
) *AdminV1DeviceService {
	l := log.NewHelper(log.With(logger, "module", "service/device"))
	return &AdminV1DeviceService{
		log:                 l,
		deviceRepo:          deviceRepo,
		deviceHeartbeatRepo: deviceHeartbeatRepo,
	}
}

type AdminV1DeviceService struct {
	pb.UnimplementedDeviceServer
	log                 *log.Helper
	deviceRepo          *data.DeviceRepo
	deviceHeartbeatRepo *data.DeviceHeartbeatRepo
}
