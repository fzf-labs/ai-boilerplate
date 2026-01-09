//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/server"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/service"
	conf "github.com/fzf-labs/kratos-contrib/api/conf/v1"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
