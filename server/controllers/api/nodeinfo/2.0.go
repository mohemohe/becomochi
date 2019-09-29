package nodeinfo

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/util"
	"net/http"
)

type (
	NodeInfo20 struct {
		Version           string   `json:"version"`
		SoftWare          SoftWare `json:"software"`
		Protocols         []string `json:"protocols"`
		Services          Services `json:"services"`
		OpenRegistrations bool     `json:"openRegistrations"`
		Usage             Usage    `json:"usage"`
		MetaData          MetaData `json:"metadata"`
	}
	SoftWare struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	Services struct {
		Inbound  []string `json:"inbound"`
		Outbound []string `json:"outbound"`
	}
	Usage struct {
		Users Users `json:"users"`
	}
	Users struct {
	}
	MetaData struct {
		Name          string         `json:"name"`
		Description   string         `json:"description"`
		Maintainer    Maintainer     `json:"maintainer"`
		Langs         []string       `json:"langs"`
		ToSUrl        string         `json:"ToSUrl"`
		RepositoryUrl string         `json:"repositoryUrl"`
		FeedbackUrl   string         `json:"feedbackUrl"`
		Announcements []Announcement `json:"announcements"`
	}
	Maintainer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	Announcement struct {
		Text  string `json:"text"`
		Image string `json:"image"`
		Title string `json:"title"`
	}
)

func GetNodeInfo20(c echo.Context) error {
	nodeInfo := NodeInfo20{
		Version: "2.0",
		SoftWare: SoftWare{
			Name:    "becomochi",
			Version: util.GetFullVersion(),
		},
		Protocols: []string{
			"activitypub",
		},
		Services: Services{
			Inbound:  []string{},
			Outbound: []string{},
		},
		OpenRegistrations: false,
		Usage: Usage{
			Users: Users{},
		},
		MetaData: MetaData{
			Name:        "",
			Description: "",
			Maintainer: Maintainer{
				Name:  "mohemohe",
				Email: "mohemohe@users.noreply.github.com",
			},
			Langs: []string{
				"ja",
			},
			ToSUrl:        "",
			RepositoryUrl: "https://github.com/mohemohe/becomochi",
			FeedbackUrl:   "https://github.com/mohemohe/becomochi/issues/new",
			Announcements: []Announcement{},
		},
	}
	return c.JSON(http.StatusOK, nodeInfo)
}
