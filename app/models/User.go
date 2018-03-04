package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/modules/orm/gorm/app"

	"github.com/iris-contrib/httpexpect"
)

var (
	DB *gorm.DB
)

type Users struct {
	UserId   int    `gorm:"default:'user_id'"`
	UserName string `gorm:"default:'user_name'"`
	Mobile   string `gorm:"default:'mobile'"`
	Password string `gorm:"default:'password'"`
	Avatar   string
	Status   string
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

func CreateUser(users Users)  {
	gormdb.DB.NewRecord(users)
	gormdb.DB.Create(&users)
}
