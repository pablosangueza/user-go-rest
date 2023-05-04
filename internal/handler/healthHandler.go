package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthAlive(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
