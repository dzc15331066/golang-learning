package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var myengine *xorm.Engine

func init() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	myengine = engine
	err = myengine.Sync(new(UserInfo))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
