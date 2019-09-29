package util

import (
	"github.com/labstack/echo/v4"
)

func GetActorID(c echo.Context, screenName string) string {
	baseUrl := BaseURL(c)
	return baseUrl + "/api/activitypub/" + screenName
}