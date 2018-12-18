package routers

import (
	"github.com/astaxie/beego"
	"testbdd/controllers"
)

func init() {
	beego.Router("/", &controllers.UserController{}, "get:Index")
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/purchase", &controllers.UserController{}, "get:Purchase")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
}
