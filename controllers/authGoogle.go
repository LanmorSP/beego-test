package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Operations about Users
type AuthGoogleController struct {
	beego.Controller
}

var conf = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  "",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/adwords",
	},
	Endpoint: google.Endpoint,
}

// @Title login
// @Description google ads oauth test
// @Success 200 {string} login success
// @router /login [get]
func (c *AuthGoogleController) Login() {

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	c.Ctx.Redirect(303, url)
}

// @Title login
// @Description google ads oauth test
// @Success 200 {string} login success
// @router /callback [get]
func (u *AuthGoogleController) Callback() {
	qs := u.Ctx.Request.URL.Query()

	if !qs.Has("code") {
		u.Data["json"] = "code is required"
		u.Ctx.Output.Status = 400
	}

	tk, _ := conf.Exchange(context.Background(), qs.Get("code"))

	rs := make(map[string]interface{})
	rs["token"] = tk

	resp, _ := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tk.AccessToken)
	defer resp.Body.Close()
	var userInfo interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)
	rs["user_info"] = userInfo

	u.Data["json"] = rs
	u.ServeJSON()
}
