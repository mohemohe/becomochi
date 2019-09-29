package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/controllers"
	"github.com/mohemohe/becomochi/server/controllers/api/activitypub"
	"github.com/mohemohe/becomochi/server/controllers/api/nodeinfo"
	"github.com/mohemohe/becomochi/server/middlewares"
	"github.com/mohemohe/echoHelper/v4"
)

func main() {
	eh := echoHelper.New(echo.New())
	eh.RegisterRoutes([]echoHelper.Route{
		{echo.GET, ".well-known/host-meta", controllers.GetWellKnownHostMeta, nil},
		{echo.GET, ".well-known/webfinger", controllers.GetWellKnownWebFinger, nil},
		{echo.GET, ".well-known/nodeinfo", controllers.GetWellKnownNodeInfo, nil},
		{echo.GET, "api/nodeinfo/2.0", nodeinfo.GetNodeInfo20, nil},
		{echo.GET, "api/activitypub/:screenName", activitypub.GetActivityPubUser, &[]echo.MiddlewareFunc{middlewares.RedirectBrowser}},
		{echo.POST, "api/activitypub/:screenName/inbox", activitypub.PostActivityPubInboxByUser, nil},
		{echo.POST, "api/activitypub/:screenName/outbox", activitypub.PostActivityPubOutboxByUser, nil},
	})
	eh.Serve()
}