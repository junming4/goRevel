package models

import (
	"github.com/revel/modules/orm/gorm/app"
	"strings"
)

type Roles struct {
	RoleId      int `gorm:"default:'role_id',primary_key"`
	Name        string
	Slug        string
	Description string
	Level       int8
	CreatedTime int32 `gorm:"default:'created_time'"`
	UpdatedAt string
}
/*
* 获取角色列表
*/
func RoleList(roleId int, name string, slug string, limitOffset string, order string) []Roles {
	var roles []Roles
	var whereRole Roles

	limits := strings.Split(limitOffset,",")

	if roleId > 0 {
		whereRole.RoleId = roleId
	}
	if len(name) > 0 {
		whereRole.Name = name
	}
	if len(slug) > 0 {
		whereRole.Slug = slug
	}
	gormdb.DB.Where(whereRole).Offset(limits[0]).Limit(limits[1]).Order(order).Find(&roles)
	return roles
}

func CreateRole(roles Roles) bool {
	if gormdb.DB.NewRecord(roles) != true {
		return false
	}
	gormdb.DB.Create(&roles)
	return true
}

func GetRole(id int) Roles {
	var roles Roles
	gormdb.DB.Where(&Roles{RoleId: id}).First(&roles)
	return roles
}