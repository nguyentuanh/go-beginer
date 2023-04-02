package main

import (
	"go-template/config"

	handleossignal "go-template/pkg/handle-os-signal"
	"go-template/pkg/l/sentry"
	tracer "go-template/pkg/trace"

	"go-template/pkg/container"
	"go-template/pkg/l"
)

// @title Documentation API
// @version 2.0
// @description Documentation API .

// @contact.name API Support
// @contact.url http://localhost
// @contact.email

// @host localhost
// @BasePath /
// @schemes https
func main() {
	ll := l.New()
	cfg := config.Load(ll)
	if cfg.SentryConfig.Enabled {
		ll = l.NewWithSentry(&sentry.Configuration{
			DSN: cfg.SentryConfig.DNS,
			Trace: struct{ Disabled bool }{
				Disabled: !cfg.SentryConfig.Trace,
			},
		})
	}
	container.Register(func() l.Logger {
		return ll
	})
	tracer.NewWithOption(
		&tracer.Option{
			TracerName: cfg.Tracing.Name + ":" + cfg.Environment,
			L:          ll,
			TelExporter: &tracer.TelExporter{
				ID:    cfg.Tracing.TeleID,
				Token: cfg.Tracing.TeleToken,
			},
			Enable: true,
		},
	)

	// init os signal handle
	shutdown := handleossignal.New(ll)
	shutdown.HandleDefer(func() {
		ll.Sync()
	})
	container.Register(func() handleossignal.IShutdownHandler {
		return shutdown
	})

	bootstrap(cfg)

	go registerHttpServer(cfg)

	// handle signal
	if cfg.Environment == "D" {
		shutdown.SetTimeout(1)
	}
	shutdown.Handle()
}
