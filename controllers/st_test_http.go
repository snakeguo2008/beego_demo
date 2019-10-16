package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type StHttpPost struct {
	beego.Controller
}

type Response struct {
	Ret uint32 `json:"ret"`
	Msg string `json:"msg"`
}

func (s *StHttpPost) Test() {
	s.Ctx.Output.Header("Content-Type","application/json;charset=utf-8")

	var resp Response
	resp.Ret = 0
	resp.Msg = "this is response message"

	info , _ := json.Marshal(resp)
	s.Ctx.WriteString(string(info))
}
