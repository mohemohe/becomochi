package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/util"
	"github.com/mohemohe/temple"
	"net/http"
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

	// TODO: resource
	screenName := "mohemohe"
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
