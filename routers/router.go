package routers

import (
	"github.com/dazhenghu/beecms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
