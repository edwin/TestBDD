package controllers

import (
	"github.com/astaxie/beego"
	"testbdd/models"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Index() {
	u.Data["json"] = "hello world"
	u.ServeJSON()
}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"

		session := u.StartSession()
		session.Set("username", username)
	} else {
		u.Data["json"] = "user not exist"
		u.Ctx.ResponseWriter.WriteHeader(404)
	}
	u.ServeJSON()
}

func (u *UserController) Purchase() {
	name := u.GetSession("username")
	if name == nil {
		u.Data["json"] = "forbidden"
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.ServeJSON()

		return
	}

	u.Data["json"] = "successfully purchase"
	u.ServeJSON()
}

func (u *UserController) Logout() {

	u.DelSession("username")
	u.Data["json"] = "successfully logout"
	u.ServeJSON()
}
