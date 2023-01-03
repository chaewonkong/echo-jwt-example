package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var AppSet = wire.NewSet(NewRepository, NewHandler, NewServer, echo.New)
