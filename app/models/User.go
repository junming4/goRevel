package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/modules/orm/gorm/app"

)

var (
	DB *gorm.DB
)

type Users struct {
	UserId   int    `gorm:"default:'user_id',primary_key"`
	UserName string `gorm:"default:'user_name'"`
	Mobile   string `gorm:"default:'mobile'"`
	Password string `gorm:"default:'password'"`
	Avatar   string
	Status   int
}

func GetUser() *Users {
	users := new(Users)
	gormdb.DB.First(&users)
	return users
}

func GetUserById(userId int) (*Users) {
	users := new(Users)
	gormdb.DB.Where(&Users{UserId: userId}).First(&users)
	return users
}
/**
*
* 创建用户数据
*/
func CreateUser(users Users) bool {
	if gormdb.DB.NewRecord(users) != true {
		return false
	}
	gormdb.DB.Create(&users)
	return true
}
