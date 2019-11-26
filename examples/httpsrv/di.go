//+build wireinject

package main

import "github.com/google/wire"

func initApp() *app{
	wire.Build(set)
	return &app{}
}
