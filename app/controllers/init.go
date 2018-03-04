package controllers

import (
	"github.com/jinzhu/gorm"
	//"github.com/revel/revel"
)

//xorm orm 操作
/*import (
	"github.com/revel/revel"
	"myapp/app/models"
	"github.com/go-xorm/xorm"
)
var(
	engine *xorm.Engine
)

func init() {
	revel.OnAppStart(Init)
}

func Init() {
	engine = models.GetEngine()
}*/


var (
	db *gorm.DB
)

func init() {
	//revel.OnAppStart(Init)
}
