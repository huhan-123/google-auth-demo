// https://github.com/hertz-contrib/jwt

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"google.golang.org/api/idtoken"
	"log"
	"net/http"
	"time"
)

const IdentityKey = "user_id"
const ClientId = "your client_id"

var AuthMiddleware *jwt.HertzJWTMiddleware

func InitAuthMiddleware() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:             "cutewallpapershub auth",
		Key:               []byte("mdfulsxbvgdgwittcdhnyxfd"),
		Timeout:           2 * 24 * time.Hour,
		MaxRefresh:        time.Hour,
		IdentityKey:       IdentityKey,
		SendAuthorization: true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				// only return the user_id to token
				return jwt.MapClaims{
					IdentityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			id := int64(claims[IdentityKey].(float64))
			return id
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			// parse credential
			loginReq := &LoginReq{}
			if err := c.BindAndValidate(&loginReq); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			payload, err := idtoken.Validate(ctx, loginReq.Credenticial, ClientId)
			if err != nil {
				log.Println(err.Error())
				return "", jwt.ErrMissingLoginValues
			}
			infoMap := payload.Claims
			name := infoMap["name"]
			email := infoMap["email"]
			avatar := infoMap["picture"]

			// select user_info by email from database
			user := GetUserByEmail(email.(string))

			// if it has user_info, then return it
			// if it has no user_info , then create info and return it
			if user == nil {
				user = &User{
					Id:       4,
					UserName: name.(string),
					Email:    email.(string),
					Avatar:   avatar.(string),
				}
				AddUser(user)
			}

			return user, nil
		},
		// return true if user have the permission to access this endpoint
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//if v, ok := data.(*User); ok && v.UserName == "admin" {
			//	return true
			//}
			//
			//return false
			return true
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": http.StatusOK,
				"msg":  "",
				"data": map[string]interface{}{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				},
			})
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": http.StatusOK,
				"msg":  "",
			})
		},
		//Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
		//	c.JSON(code, map[string]interface{}{
		//		"code":    code,
		//		"message": message,
		//	})
		//},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer". If you want empty value, use WithoutDefaultTokenHeadName.
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := AuthMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type LoginReq struct {
	Credenticial string `json:"credential"`
}

// pretend to be a database
var Users = []*User{
	{
		Id:       1,
		UserName: "Ethan Walker",
		Email:    "111111111@gmail.com",
		Avatar:   "https://www.google.com",
	},
	{
		Id:       2,
		UserName: "Olivia Carter",
		Email:    "2222222222@gmail.com",
		Avatar:   "https://www.google.com",
	}, {
		Id:       3,
		UserName: "Mason Davis",
		Email:    "3333333333@gmail.com",
		Avatar:   "https://www.google.com",
	},
}

func GetUserById(id int64) *User {
	for _, user := range Users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func GetUserByEmail(email string) *User {
	for _, user := range Users {
		if user.Email == email {
			return user
		}
	}
	return nil
}

func AddUser(user *User) {
	Users = append(Users, user)
}
