package main

import (
	"flag"
	"os"

	"github.com/fzf-labs/kratos-contrib/bootstrap"
	"github.com/fzf-labs/kratos-contrib/pkg/mq"
	_ "go.uber.org/automaxprocs"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// id is the hostname of the compiled software.
	id, _ = os.Hostname()
	// Service is the service instance.
	Service = bootstrap.NewService(
		Name,
		Version,
		id,
		map[string]string{},
	)
)

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, mqServer mq.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.ID),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
			mqServer,
		),
	)
}

func main() {
	flag.Parse()
	cfg, logger, _, _ := bootstrap.Bootstrap(Service)
	app, cleanup, err := wireApp(cfg, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
