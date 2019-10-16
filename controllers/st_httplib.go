package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
)

type StHttpLib struct {
	beego.Controller
}

func (c *StHttpLib) Get() {
	beego.Info("test beegon http lib ....")

	c.Ctx.Output.Header("Content-Type","application/json;charset=utf-8")
	info := httplib.Get("https://www.baidu.com/")
	strInfo, err := info.String()
	if err != nil{
		beego.Error("get baidu info failed")
		return
	}
	type Response struct {
		Ret uint32 `json:"ret"`
		Msg string `json:"msg"`
	}
	var resp Response
	resp.Ret = 0
	resp.Msg = strInfo

	respJson , _ := json.Marshal(resp)
	//c.Ctx.WriteString(strInfo)
	c.Ctx.WriteString(string(respJson))
}
