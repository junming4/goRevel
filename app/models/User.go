package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/modules/orm/gorm/app"
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"fmt"
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

func GetUserByUser(userName string, password string) (*Users) {
	users := new(Users)
	gormdb.DB.Where(Users{UserName:userName}).First(&users)
	fmt.Println(users)
	if users == nil {
		gormdb.DB.Where(Users{Mobile:userName}).First(&users)
		if users == nil {
			return nil
		}
	}
	revel.AppLog.Error("user")


	passwords := []byte(password)
	hashedPassword := []byte(users.Password)
	// Hashing the password with the default cost of 10
	/*hashedPassword, err := bcrypt.GenerateFromPassword(passwords, bcrypt.DefaultCost)
	if err != nil {
		return nil
	}*/
	err2 := bcrypt.CompareHashAndPassword(hashedPassword, passwords)
	if err2 != nil {
		return nil
	}
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
