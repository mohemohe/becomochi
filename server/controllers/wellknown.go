package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/util"
	"github.com/mohemohe/temple"
	"net/http"
	"strings"
)

type (
	WebFingerAccount struct {
		Subject string `json:"subject"`
		Links   []Link `json:"links"`
	}
	Link struct {
		Rel      string `json:"rel"`
		Type     string `json:"type,omitempty"`
		Href     string `json:"href,omitempty"`
		Template string `json:"template,omitempty"`
	}
	NodeInfo struct {
		Links   []Link `json:"links"`
	}
)

func GetWellKnownHostMeta(c echo.Context) error {
	xmlBase := `<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <Link rel="lrdd" type="application/xrd+xml" template="{{.baseUrl}}/.well-known/webfinger?resource={uri}"/>
</XRD>`
	xml, err := temple.Execute(xmlBase, map[string]interface{}{
		"baseUrl": util.BaseURL(c),
	})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.XMLBlob(http.StatusOK, []byte(xml))
}

func GetWellKnownWebFinger(c echo.Context) error {
	baseuUrl := util.BaseURL(c)
	resource := c.QueryParam("resource")
	if resource == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	acctURI := strings.TrimPrefix(resource, "acct:")
	screenName := "mohemohe"
	if acctURI != screenName + "@" + c.Request().Host {
		return c.NoContent(404)
	}

	acct := WebFingerAccount{
		Subject: "acct:" + screenName + "@" + c.Request().Host,
		Links: []Link{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: baseuUrl + "/api/activitypub/" + screenName,
			},
			// {
			// 	Rel:  "http://webfinger.net/rel/profile-page",
			// 	Type: "text/html",
			// 	Href: baseuUrl + "/@" + screenName,
			// },
			// {
			// 	Rel:  "http://ostatus.org/schema/1.0/subscribe",
			// 	Template: baseuUrl + "/api/activitypub/{uri}/subscribe",
			// },
		},
	}
	return c.JSON(200, acct)
}

func GetWellKnownNodeInfo(c echo.Context) error {
	baseuUrl := util.BaseURL(c)
	nodeInfo := NodeInfo{
		Links: []Link{
			{
				Rel: "http://nodeinfo.diaspora.software/ns/schema/2.0",
				Href: util.JoinURL(baseuUrl, "/api/nodeinfo/2.0"),
			},
		},
	}

	return c.JSON(http.StatusOK, nodeInfo)
}