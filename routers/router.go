package routers

import (
	"github.com/cqspirit/datasync/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
