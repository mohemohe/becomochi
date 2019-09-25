package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/controllers"
	v1 "github.com/mohemohe/becomochi/server/controllers/api/v1"
	"github.com/mohemohe/echoHelper/v4"
)

func main() {
	e := echo.New()
	eh := echoHelper.New(e)
	eh.RegisterRoutes([]echoHelper.Route{
		{echo.GET, ".well-known/host-meta", controllers.GetWellKnownHostMeta, nil},
		{echo.GET, ".well-known/webfinger", controllers.GetWellKnownWebFinger, nil},
		{echo.GET, "api/activitypub/:screenName", v1.GetActivityPubUser, nil},
	})
	eh.Serve()
}