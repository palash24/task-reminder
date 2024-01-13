package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/palash24/task-reminder/config"
	internal "github.com/palash24/task-reminder/internal"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(internal.NewHttpServer),
		fx.Provide(internal.NewDb),
		fx.Invoke(startServer),
	).Run()
}

func startServer(
	lifecycle fx.Lifecycle, srv internal.Server,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go http.ListenAndServe(fmt.Sprintf("%s:%s", srv.ServerAddr, srv.ServerPort), srv.Router)
				return nil
			},
		},
	)
}
