//go:build wireinject
// +build wireinject

package modules

import (
	"go-jwt/app"

	"github.com/google/wire"
)

func InitializeServer() app.Server {
	wire.Build(AppModules)
	return app.Server{}
}
