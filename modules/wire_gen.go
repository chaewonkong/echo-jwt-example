// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package modules

import (
	"github.com/labstack/echo/v4"
	"go-jwt/app"
)

// Injectors from wire.go:

func InitializeServer() (app.Server, error) {
	repository := app.NewRepository()
	handler := app.NewHandler(repository)
	echoEcho := echo.New()
	server := app.NewServer(handler, echoEcho)
	return server, nil
}
