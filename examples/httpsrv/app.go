package main

import (
	"net/http"

	"github.com/google/wire"
	"github.com/wingyplus/wireapp"
)

var set = wire.NewSet(
	wireapp.Set,
	wireapp.ProvideHTTPAppConfigFromEnv,
	newApp,
	newPage,
	wire.Bind(new(http.Handler), new(*page)),
)

type app struct {
	*wireapp.HTTPApp
}

func newApp(httpapp *wireapp.HTTPApp, h http.Handler) *app {
	a := &app{
		HTTPApp: httpapp,
	}
	a.Register(h)
	return a
}
