package controllers

import (
	"encoding/json"
	"test-api/models"

	"github.com/beego/beego/v2/server/web/pagination"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	uid := models.AddUser(user)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *UserController) GetAll() {
	//initial result
	rs := make(map[string]interface{})

	//count total
	totalCnt, _ := models.GetAllUsersCount()

	//pagination generator
	per := 3
	pgr := pagination.SetPaginator(c.Ctx, per, totalCnt)

	users, _ := models.GetAllUsers(pgr.Offset(), per)
	rs["data"] = users
	rs["meta"] = map[string]interface{}{
		"currentPage": pgr.Page(),
		"perPage":     pgr.PerPageNums,
		"totalCount":  pgr.Nums(),
		"totalPage":   pgr.PageNums(),
	}

	c.Data["json"] = rs

	c.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (c *UserController) Get() {
	id, err := c.GetInt(":uid")
	if err == nil {
		users, err := models.GetUser(id)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.Output.Status = 404
		} else {
			c.Data["json"] = users
		}
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	id, err := u.GetInt(":uid")
	if err == nil {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(id, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	id, err := u.GetInt(":uid")
	if err == nil {
		msg, err := models.DeleteUser(id)
		if err != nil {
			u.Data["json"] = err.Error()
			u.Ctx.Output.Status = 404
		} else {
			u.Data["json"] = msg
		}
		u.ServeJSON()
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
