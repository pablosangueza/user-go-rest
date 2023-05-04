package handler

import (
	"net/http"
	"time"

	"github.com/dom/user/internal/config"
	"github.com/dom/user/internal/users"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type UserCredential struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type AuthUserHandler struct {
	query users.UserQuery
}
type AuthUserHandlerParams struct {
	Query users.UserQuery
}

func NewAuthUserHandler(p *AuthUserHandlerParams) *AuthUserHandler {
	return &AuthUserHandler{
		query: p.Query,
	}
}

func (this AuthUserHandler) Handle(c echo.Context) error {

	credetials := new(UserCredential)
	if err := c.Bind(credetials); err != nil {
		return err
	}

	if credetials.Email != "myusername" || credetials.Password != "mypassword" {
		return echo.ErrUnauthorized
	}

	expTime := time.Now().Add(time.Hour)

	claims := jwt.MapClaims{
		"username": credetials.Email,
		"exp":      expTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.DefaultJWTConfig().SecretKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": signedToken,
	})

}
