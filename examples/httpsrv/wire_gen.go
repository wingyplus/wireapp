// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/wingyplus/wireapp"
)

// Injectors from di.go:

func initApp() *app {
	httpAppConfig := wireapp.ProvideHTTPAppConfigFromEnv()
	httpApp := wireapp.NewHTTPApp(httpAppConfig)
	mainPage := newPage()
	mainApp := newApp(httpApp, mainPage)
	return mainApp
}
