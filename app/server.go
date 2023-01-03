package app

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Server struct {
	handler Handler
	echo    *echo.Echo
}

func NewServer(h Handler, e *echo.Echo) Server {
	s := Server{h, e}

	return s
}

func NewClaimsFunc(c echo.Context) jwt.Claims {
	return new(JwtCustomClaims)
}

func (s Server) Start() {
	s.bindRoute()
	s.echo.Logger.Fatal(s.echo.Start(":8080"))
}

func (s Server) bindRoute() {
	s.echo.POST("/login", s.handler.login)

	r := s.echo.Group("/auth")

	config := echojwt.Config{
		NewClaimsFunc: NewClaimsFunc,
		SigningKey:    []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", s.handler.restricted)
}
