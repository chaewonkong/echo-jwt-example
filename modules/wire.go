//go:build wireinject
// +build wireinject

package modules

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeServer() Server {
	wire.Build(NewRepository, NewHandler, echo.New, NewServer)
	return Server{}
}
