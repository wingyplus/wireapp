package wireapp

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
)

var Set = wire.NewSet(
	NewHTTPApp,
)

// App provides specification for implementing application shell.
type App interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
	Run()
}

// HTTPAppConfig is configuration for HTTPApp.
type HTTPAppConfig struct {
	Addr string `envconfig:"ADDR"`

	TLSCertFile string `envconfig:"TLS_CERT_FILE"`
	TLSKeyFile  string `envconfig:"TLS_KEY_FILE"`
	TLSEnabled  bool   `envconfig:"TLS_ENABLED"`
}

func ProvideHTTPAppConfigFromEnv() HTTPAppConfig {
	var cfg HTTPAppConfig
	envconfig.MustProcess("HTTP", &cfg)
	return cfg
}

// HTTPApp provides application that uses HTTP base protocol.
type HTTPApp struct {
	server *http.Server
	cfg    HTTPAppConfig
}

func (app *HTTPApp) Start(ctx context.Context) error {
	if len(app.cfg.Addr) == 0 {
		app.server.Addr = ":8080"
	} else {
		app.server.Addr = app.cfg.Addr
	}
	if app.cfg.TLSEnabled {
		return app.server.ListenAndServeTLS(app.cfg.TLSCertFile, app.cfg.TLSKeyFile)
	}
	return app.server.ListenAndServe()
}

func (app *HTTPApp) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}

func (app *HTTPApp) Run() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go app.Start(context.TODO())

	<-sig
	app.Shutdown(context.TODO())
}

// Register handler into http server.
func (app *HTTPApp) Register(handler http.Handler) {
	app.server.Handler = handler
}

// NewHTTPApp creating http application.
func NewHTTPApp(cfg HTTPAppConfig) *HTTPApp {
	return &HTTPApp{
		server: &http.Server{},
		cfg:    cfg,
	}
}
