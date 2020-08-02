package routers

import (
	"github.com/alex-levin/mosaic-b/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
