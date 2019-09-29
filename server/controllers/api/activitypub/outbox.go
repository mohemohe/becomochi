package activitypub

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PostActivityPubOutboxByUser(c echo.Context) error {
	return c.NoContent(http.StatusInternalServerError)
}