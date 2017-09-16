package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/logs"
)

func init() {
	orm.Debug = true
	db := beego.AppConfig.String("db")
	user := beego.AppConfig.String("dbUser")
	password := beego.AppConfig.String("dbPassword")
	host := beego.AppConfig.String("dbHost")
	dbSource := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s sslmode=disable",
		user, password, host, db)
	logs.Debug("db dbSource: %s", dbSource)
	orm.RegisterDataBase(
		"default", "postgres",
		dbSource,
	)
}
