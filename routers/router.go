package routers

import (
	"github.com/devbus/devbus/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/app", "webapp/dist")
	beego.SetStaticPath("/static", "webapp/dist/static")

	ns := beego.NewNamespace("/api",
		beego.NSRouter("/session/login", &controllers.LoginController{}),
		beego.NSRouter("/session/register", &controllers.RegisterController{}, "post:Register"),
		beego.NSRouter("/session/activate", &controllers.RegisterController{}, "get:Activate"),

		beego.NSRouter("/project", &controllers.ProjectController{}),
		beego.NSRouter("/issue", &controllers.IssueController{}),
	)
	beego.AddNamespace(ns)
}
