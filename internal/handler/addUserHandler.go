package handler

import (
	"net/http"

	"github.com/dom/user/internal/users"
	"github.com/labstack/echo"
)

type RequestUser struct {
	UserName string `json: "userName"`
	LastName string `json: "lastName"`
	Email    string `json: "email"`
	Role     string `json: "role"`
}

type ResponseNewUser struct {
	UserId int `json: "userId"`
}

type AddUserHandler struct {
	cmd users.UserCommand
}

type AddUserHandlerParams struct {
	Cmd users.UserCommand
}

func NewAddUserHandler(p *AddUserHandlerParams) *AddUserHandler {
	return &AddUserHandler{
		cmd: p.Cmd,
	}
}

func (this AddUserHandler) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	requestUser := new(RequestUser)
	if err := c.Bind(requestUser); err != nil {
		return err
	}

	res, err := this.cmd.SaveUser(ctx, users.User{
		UserName: requestUser.UserName,
		LastName: requestUser.LastName,
		Email:    requestUser.Email,
		Role:     requestUser.Role,
	})
	if err != nil {
		return err
	}

	response := &ResponseNewUser{
		UserId: res,
	}

	return c.JSON(http.StatusOK, response)

}
