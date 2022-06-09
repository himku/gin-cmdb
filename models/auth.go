package models

import (
	"GoNotes/middleware"
	"gorm.io/gorm"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/8 10:19
 **/

// Users 用户表结构体数据
type Users struct {
	gorm.Model
	ID       int64  `gorm:"primary_key"` // 主键ID
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

var db = middleware.NewConnectMySQL()
var user Users

func init() {
	// 初始化表
	err := db.AutoMigrate(&user)
	if err != nil {
		return
	}
}

func CreateUser(user *Users) bool {

	if err := db.Debug().Save(&user); err != nil {
		return false
	} else {
		return true
	}
}
