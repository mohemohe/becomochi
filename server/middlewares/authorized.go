package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/configs"
	"net/http"
)

func Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("X-BECOMOCHI-AUTH")
		if auth == configs.GetEnv().Becomochi.Api.Key {
			return next(c)
		}
		return c.NoContent(http.StatusUnauthorized)
	}
}