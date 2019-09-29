package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/util"
	"net/http"
	"strings"
)

func RedirectBrowser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accept := c.Request().Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			path := c.Request().URL.Path
			to := strings.Trim(path, "api/activitypub/")
			baseUrl := util.BaseURL(c)
			return c.Redirect(http.StatusFound, util.JoinURL(baseUrl, to))
		} else {
			return next(c)
		}
	}
}