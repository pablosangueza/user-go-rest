package handler

import (
	"net/http"
	"strconv"

	"github.com/dom/user/internal/users"
	"github.com/labstack/echo"
)

type GetUsersHandlerParams struct {
	Query users.UserQuery
}

type GetUsersHandler struct {
	query users.UserQuery
}

func NewGetUsersHandler(p *GetUsersHandlerParams) *GetUsersHandler {
	return &GetUsersHandler{
		query: p.Query,
	}
}

func (s GetUsersHandler) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("id")
	if idStr == "" {
		res, err := s.query.GetUsers(ctx, int32(0))
		if err != nil {
			return nil
		}
		return c.JSON(http.StatusOK, res)

	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}

		res, err := s.query.GetUsers(ctx, int32(id))
		if err != nil {
			return nil
		}
		return c.JSON(http.StatusOK, res)

	}

}
