package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"fmt"
)


var (
	db *gorm.DB
)


func init() {
	revel.TemplateFuncs["can"] = func(urlName string, request interface{}) bool {
		revel.AppLog.Error(urlName)
		fmt.Println(request)
		return false
	}
}
