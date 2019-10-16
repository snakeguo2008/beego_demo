package main

import (
	_ "beego_demo/routers"
	"github.com/astaxie/beego"
	"beego_demo/models"
)

func main() {
	models.InitMysql()
	beego.Run()
}

