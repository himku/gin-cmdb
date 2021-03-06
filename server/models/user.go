package models

import (
	"gin-cmdb/server/utils"
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

var auth Users

func CreateUser(user *Users) bool {
	if CheckUser(user.Username) {
		return false
	} else {
		db.Debug().Create(&user)
		return true
	}
}

func CheckUser(username string) bool {
	checkResult := db.Debug().Select("id").Where(Users{Username: username}).First(&auth)
	if checkResult.RowsAffected == 0 {
		return false
	}
	return true
}

func CheckAuth(username, password string) bool {
	db.Debug().Select("password").Where(Users{Username: username}).First(&auth)
	if utils.CheckPasswordHash(password, auth.Password) {
		return true
	}
	return false
}

// GetUserList 获取用户列表
func GetUserList(page int, pageSize int) (int, []interface{}) {
	var users []Users
	userList := make([]interface{}, 0, len(users))
	offset := (page - 1) * pageSize
	result := db.Debug().Offset(offset).Limit(pageSize).Find(&users)
	if result.RowsAffected == 0 {
		return 0, userList
	}
	total := len(users)
	result.Offset(offset).Limit(pageSize).Find(&users)
	for _, userSingle := range users {
		userMap := map[string]interface{}{
			"id":       userSingle.ID,
			"username": userSingle.Username,
		}
		userList = append(userList, userMap)
	}
	return total, userList
}

func DeleteUser(username string) bool {
	deleteUser := Users{Username: username}
	if CheckUser(username) {
		db.Debug().Unscoped().Where("username = ?", username).Delete(&deleteUser)
		return true
	}
	return false
}

func EditUser(u *Users) bool {
	if CheckUser(u.Username) {
		db.Debug().Model(&auth).Updates(&u)
		return true
	}
	return false
}
