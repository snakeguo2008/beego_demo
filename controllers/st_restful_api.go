package controllers

import (
	"beego_demo/comm"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"github.com/cihub/seelog"
)

type StRestfulApI struct {
	beego.Controller
}

func (s *StRestfulApI) GetName() {
	//定义http返回协议头格式为json
	s.Ctx.Output.Header("Content-Type", "application/json;charset=utf-8")
	ormOBJ := orm.NewOrm()
	sqlRaw := fmt.Sprintf("select user_id from `user` limit 10")

	var userList []uint64
	ormOBJ.Raw(sqlRaw).QueryRows(&userList)

	var rsp comm.StResp
	rsp.Ret = 0
	rsp.ErrMsg = "success"
	rsp.UserList = userList
	seelog.Infof("GetName:%v", rsp)
	data, _ := json.Marshal(rsp)
	s.Ctx.WriteString(string(data))
}
