package models

import "GoNotes/middleware"

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/8 10:19
 **/

type User struct {
	ID       int    `gorm:"primary_key" json:"id"` // 主键ID
	Username string `json:"username"`
	Password string `json:"password"`
}

var db = middleware.NewConnectMySQL()

func SelectUser(user User) {

}
