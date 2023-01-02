package modules

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	repository Repository
}

func NewHandler(r Repository) Handler {
	return Handler{r}
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (h Handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}

	// create user
	h.repository.createUser(username, password)

	claims := &JwtCustomClaims{
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

func (h Handler) restricted(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*JwtCustomClaims)
	name := claims.Name

	user, err := h.repository.findUser(name)

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Welcome "+user.Username+"!")
}
