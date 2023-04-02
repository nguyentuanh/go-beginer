package main

import (
	"go-template/config"
	api "go-template/internal/app-api"
	"go-template/internal/app-api/handler/health"
	"go-template/pkg/container"
	handleossignal "go-template/pkg/handle-os-signal"
	"go-template/pkg/l"
)

func registerHttpServer(cfg *config.Config) {
	var ll = container.ResolverMust[l.Logger]()
	var healthHandler = container.ResolverMust[health.Controller]()
	var shutdown = container.ResolverMust[handleossignal.IShutdownHandler]()

	healthHandler.SetReady(true)
	shutdown.HandleDefer(func() {
		healthHandler.SetReady(false)
	})

	gw := api.New(cfg.Environment)
	gw.Middleware(ll)
	gw.InitHealth(healthHandler)
	gw.InitLogHandler()
	gw.InitMetrics()
	gw.InitHome()
	ll.Info("HTTP server start listening", l.Any("HTTP address", cfg.GetHTTPAddress()))
	err := gw.Listen(cfg.GetHTTPAddress())
	if err != nil {
		ll.Fatal("error listening to address", l.Any("address", cfg.GetHTTPAddress()), l.Error(err))
		return
	}
}
