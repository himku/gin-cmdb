package models

import (
	"gin-cmdb/middleware"
	"gin-cmdb/utils"
	"github.com/jinzhu/gorm"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/8 10:19
 **/

// Users 用户表结构体数据
type Users struct {
	gorm.Model
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

var db = middleware.NewConnectMySQL()
var auth Users

func init() {
	// 初始化表
	err := db.AutoMigrate(&Users{})
	if err != nil {
		return
	}
}

func CreateUser(user *Users) bool {
	if CheckUser(user.Username) {
		return false
	} else {
		db.Debug().Create(&user)
		return true
	}
}

func CheckUser(username string) bool {
	db.Debug().Select("id").Where(Users{Username: username}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func CheckAuth(username, password string) bool {
	db.Debug().Select("password").Where(Users{Username: username}).First(&auth)
	if utils.CheckPasswordHash(password, auth.Password) {
		return true
	}
	//	if auth.ID > 0 {
	//		return true
	//	}
	//	return false
	//}
	return false
}
