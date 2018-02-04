package models

import (
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"fmt"
)

var (
	engine *xorm.Engine
)

func init()  {

}

func InitDb()  {
	var err error
	driverName := revel.Config.StringDefault("db.driver", "mysql")
	dbname,_ := revel.Config.String("dbname")
	user, _ := revel.Config.String("db.user")
	password, _ := revel.Config.String("db.password")
	host, _ := revel.Config.String("db.host")
	engine,err = xorm.NewEngine(driverName, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname))
	if err != nil {
		revel.ERROR.Panicln(err)
	}
}
