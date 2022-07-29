package controllers

import (
	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type TestLogController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /send [get]
func (c *TestLogController) Logger() {
	logs.Warning("hello")

	c.Data["json"] = "heelo"

	c.ServeJSON()
}
