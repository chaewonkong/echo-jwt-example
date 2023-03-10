//go:build wireinject
// +build wireinject

package modules

import (
	"go-jwt/app"

	"github.com/google/wire"
)

func InitializeServer() (app.Server, error) {
	wire.Build(AppModules)
	return app.Server{}, nil
}
