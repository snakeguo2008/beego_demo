package models

import (
	"beego_demo/comm"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() {
	seelog.Infof("this init models \n")
	beego.AppConfig.String("mysql")
	mysqlConf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local", beego.AppConfig.String("mysql_user"), beego.AppConfig.String("mysql_pwd"),
		beego.AppConfig.String("mysql_host"), beego.AppConfig.String("mysql_port"), beego.AppConfig.String("mysql_db"))

	MysqlMaxIdle, _ := beego.AppConfig.Int("mysql_maxidle")
	MysqlMaxConn, _ := beego.AppConfig.Int("mysql_maxconn")


	//beego.Info("init mysql cnf:", mysqlConf)
	fmt.Println("init mysql cnf:", mysqlConf)
	orm.RegisterDataBase("default", "mysql", mysqlConf)

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.SetMaxIdleConns("default", MysqlMaxIdle)
	orm.SetMaxOpenConns("default", MysqlMaxConn)

	orm.RegisterModel(new(comm.User))

}
