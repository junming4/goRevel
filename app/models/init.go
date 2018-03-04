package models

import (
	"github.com/revel/revel"
	"github.com/revel/modules/orm/gorm/app"
)


func init()  {
	revel.OnAppStart(gormdb.InitDB)
}
