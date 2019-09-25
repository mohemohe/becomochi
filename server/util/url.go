package util

import (
	"github.com/labstack/echo/v4"
)

func BaseURL(c echo.Context) string {
	proto := "http://"
	if c.Request().TLS != nil {
		proto = "https://"
	}
	host := c.Request().Host

	return proto + host
}