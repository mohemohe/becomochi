package activitypub

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/models"
	"net/http"
)

func GetActivityPubUser(c echo.Context) error {
	if c.Param("screenName") != "mohemohe" {
		return c.NoContent(http.StatusNotFound)
	}

	user := models.User{
		Email:        "いや〜ん",
		Password:     "あは〜ん",
		ScreenName:   "mohemohe",
		DisplayName:  "うんちになてり",
		Summary:      "またまともににAP喋れないからフォローはやめとけ",
		IconURL:      "",
	}
	return c.JSON(http.StatusOK, user.ToActivityPubPerson(c))
}