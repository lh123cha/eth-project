package routers

import (
	"github.com/astaxie/beego/plugins/cors"
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

    beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.LoginController{},"post:login")
	beego.Router("/add", &controllers.AddController{},"Post:Add")
	beego.Router("/addorder", &controllers.AddorderController{},"Post:Addorder")

	beego.Router("/output", &controllers.OutController{},"Post:Out")
	beego.Router("/myself", &controllers.MyController{},"Post:Myinfo")
}
