package modules

import (
	"go-jwt/app"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var AppModules = wire.NewSet(app.NewRepository, app.NewHandler, app.NewServer, echo.New)
