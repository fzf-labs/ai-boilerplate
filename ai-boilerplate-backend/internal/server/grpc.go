package server

import (
	v1 "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/service"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/fzf-labs/kratos-contrib/bootstrap"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Bootstrap,
	logger log.Logger,
	adminV1DeviceService *service.AdminV1DeviceService,
) *grpc.Server {
	srv := bootstrap.NewGrpcServer(c, logger)
	v1.RegisterDeviceServer(srv, adminV1DeviceService)
	return srv
}
