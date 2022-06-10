package middleware

import (
	"GoNotes/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/6 11:40
 **/

func NewConnectMySQL() *gorm.DB {
	c := config.NewConfig()
	// 修复数据入库时间的问题
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", c.MySQL.UserName, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.DbName)
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("connect mysql error detail %s", err)
	}
	return db
}
