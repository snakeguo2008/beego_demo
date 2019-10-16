package controllers

import "github.com/astaxie/beego"

type TplController struct {
	beego.Controller
}

func (s *TplController) TestTpl() {
	s.TplName = "test.tpl"
}
