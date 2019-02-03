package routers

import (
	"go-beego-sample/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/list", &controllers.ListController{})
		beego.Router("/create", &controllers.CreateController{})

}
