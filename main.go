package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Username string
	Password string
}

type handler struct {
	db *gorm.DB
}

func (h handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}

	// save user
	h.db.Create(&User{
		Username: username,
		Password: password,
	})

	claims := &jwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (h handler) restricted(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*jwtCustomClaims)
	name := claims.Name

	var user User
	h.db.First(&user, "username = ?", name)

	return c.String(http.StatusOK, "Welcome "+user.Username+"!")
}

func main() {
	// gorm
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!~!")
	}

	// create table
	db.AutoMigrate(&User{})

	e := echo.New()
	h := handler{
		db,
	}

	e.POST("/login", h.login)

	restrictedRoute := e.Group("/auth")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	restrictedRoute.Use(echojwt.WithConfig(config))
	restrictedRoute.GET("", h.restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
