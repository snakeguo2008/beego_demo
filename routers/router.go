package routers

import (
	"beego_demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/httplib", &controllers.StHttpLib{})
	beego.Router("/http_post", &controllers.StHttpPost{},"post:Test")
	beego.Router("/http_test_tpl",&controllers.TplController{},"post:TestTpl")
	beego.Router("/api/name",&controllers.StRestfulApI{},"get:GetName")
}

