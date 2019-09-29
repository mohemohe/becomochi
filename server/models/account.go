package models

import (
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/becomochi/server/util"
)

type (
	User struct {
		bongo.DocumentBase `bson:",inline"`
		Email              string `bson:"email" json:"email"`
		Password           string `bson:"password" json:"-"`
		ScreenName         string `bson:"screen_name" json:"screen_name"`
		DisplayName        string `bson:"display_name" json:"display_name"`
		Summary            string `bson:"summary" json:"summary"`
		IconURL            string `bson:"icon_url" json:"-"`
	}

	ActivityPubPerson struct {
		AtContext   []interface{}              `json:"@context"`
		ID          string                     `json:"id"`
		Type        string                     `json:"type"`
		ScreenName  string                     `json:"preferredUsername"`
		DisplayName string                     `json:"name"`
		Summary     string                     `json:"summary"`
		InboxURL    string                     `json:"inbox"`
		OutboxURL   string                     `json:"outbox"`
		Endpoints   ActivityPubPersonEndpoints `json:"endpoints"`
		ProfileURL  string                     `json:"url"`
		Icon        ActivityPubPersonIcon      `json:"icon"`
	}

	ActivityPubPersonIcon struct {
		Type string `json:"type"`
		Mime string `json:"mediaType"`
		URL  string `json:"url"`
	}

	ActivityPubPersonEndpoints struct {
		SharedInboxUrl string `json:"sharedInbox"`
	}

	// Users struct {
	// 	Info  *bongo.PaginationInfo `bson:"-" json:"info"`
	// 	Users []User                `bson:"-" json:"users"`
	// }
	//
	// JwtClaims struct {
	// 	ID    string `json:"id"`
	// 	Email string `json:"email"`
	// 	Role  int    `json:"role"`
	// 	jwt.StandardClaims
	// }
)

func (this *User) ToActivityPubPerson(c echo.Context) ActivityPubPerson {
	baseUrl := util.BaseURL(c)
	return ActivityPubPerson{
		AtContext: []interface{}{
			"https://www.w3.org/ns/activitystreams",
		},
		ID:          util.GetActorID(c, this.ScreenName),
		Type:        "Person",
		ScreenName:  this.ScreenName,
		DisplayName: this.DisplayName,
		Summary:     this.Summary,
		InboxURL:    baseUrl + "/api/activitypub/" + this.ScreenName + "/inbox",
		OutboxURL:   baseUrl + "/api/activitypub/" + this.ScreenName + "/outbox",
		Endpoints: ActivityPubPersonEndpoints{
			SharedInboxUrl: baseUrl + "/api/activitypub/_/inbox",
		},
		Icon: ActivityPubPersonIcon{
			Type: "Image",
			Mime: "image/png",
			URL:  "https://media.mstdn.plusminus.io/accounts/avatars/000/000/001/original/b02a7a0d5d1646b0970ac0cd6396cd90.png",
		},
	}
}

func (this *User) ToUserID() string {
	return ""
}

// const (
// 	RootRole = iota + 1
// 	UserRole
// )
//
// const (
// 	Email = "email"
// )
//
// func GetUserById(id string) *User {
// 	conn := connection.Mongo()
//
// 	user := &User{}
// 	err := conn.Collection(collections.Users).FindById(bson.ObjectIdHex(id), user)
// 	if err != nil {
// 		return nil
// 	}
//
// 	return user
// }
//
// func GetUserByEmail(email string) *User {
// 	conn := connection.Mongo()
//
// 	user := &User{}
// 	err := conn.Collection(collections.Users).FindOne(bson.M{
// 		Email: email,
// 	}, user)
// 	if err != nil {
// 		return nil
// 	}
//
// 	return user
// }
//
// func GetUsers(perPage int, page int) *Users {
// 	conn := connection.Mongo()
//
// 	result := conn.Collection(collections.Users).Find(bson.M{})
// 	if result == nil {
// 		return nil
// 	}
// 	info, err := result.Paginate(perPage, page)
// 	if err != nil {
// 		return nil
// 	}
// 	users := make([]User, info.RecordsOnPage)
// 	for i := 0; i < info.RecordsOnPage; i++ {
// 		_ = result.Next(&users[i])
// 	}
//
// 	return &Users{
// 		Info:  info,
// 		Users: users,
// 	}
// }
//
// func UpsertUser(user *User) error {
// 	if !util.IsBcrypt(user.Password) {
// 		user.Password = *util.Bcrypt(user.Password)
// 	}
// 	return connection.Mongo().Collection(collections.Users).Save(user)
// }
//
// func DeleteUser(user *User) error {
// 	return connection.Mongo().Collection(collections.Users).DeleteDocument(user)
// }
//
// func AuthroizeUser(email string, password string) (*User, *string) {
// 	user := GetUserByEmail(email)
// 	if user == nil {
// 		panic("user not found")
// 	}
//
// 	if !util.CompareHash(password, user.Password) {
// 		panic("wrong password")
// 	}
//
// 	claims := &JwtClaims{
// 		user.GetId().Hex(),
// 		user.Email,
// 		user.Role,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 		},
// 	}
//
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ts, err := token.SignedString([]byte(configs.GetEnv().Sign.Secret))
// 	if err != nil {
// 		panic("couldnt create token")
// 	}
//
// 	return user, &ts
// }
