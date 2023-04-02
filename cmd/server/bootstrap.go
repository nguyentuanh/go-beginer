package main

import (
	"context"

	"go-template/config"
	"go-template/internal/app-api/handler/health"
	"go-template/pkg/container"
	handleossignal "go-template/pkg/handle-os-signal"
)

func bootstrap(cfg *config.Config) {
	// var ll = container.ResolverMust[l.Logger]()
	var shutdown = container.ResolverMust[handleossignal.IShutdownHandler]()

	container.Register(func() health.Controller {
		return health.New()
	})

	_, cancel := context.WithCancel(context.Background())
	shutdown.HandleDefer(cancel)

	// region register Repository - db
	// - init db connection
	//endregion

	// region register Repository - redis
	// endregion

	// region register Repository - agency
	// endregion

	// region even bus
	// endregion

	// region register factory
	// endregion

	// region register publisher
	// endregion

	// region register service to container
	// endregion
}
