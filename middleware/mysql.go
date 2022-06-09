package middleware

import (
	"GoNotes/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/6 11:40
 **/

func NewConnectMySQL() *gorm.DB {
	c := config.NewConfig()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.MySQL.UserName, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.DbName)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	db.Debug()
	if err != nil {
		fmt.Printf("connect mysql error detail %s", err)
	}
	return db
}
