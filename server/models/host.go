package models

import "github.com/jinzhu/gorm"

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/17 17:47
 **/

type Host struct {
	gorm.Model
	Hostname string `json:"hostname" gorm:"commit:'主机名'"`
	OS       string `json:"os" gorm:"commit:'操作系统类型'"`
	CpuCores int    `json:"cpuCores" gorm:"commit:'CPU核心数'"`
	Status   bool   `json:"status" gorm:"commit:'主机状态'"`
}

func GetHostList(page int, pageSize int) (int, []interface{}) {
	hosts := make([]Host, 10)
	total := len(hosts)
	hostList := make([]interface{}, 0, total)
	offset := (page - 1) * pageSize
	result := db.Debug().Offset(offset).Limit(pageSize).Find(&hosts)
	if result.RowsAffected == 0 {
		return 0, hostList
	}
	result.Offset(offset).Limit(pageSize).Find(&hosts)
	for _, hostSingle := range hosts {
		hostMap := map[string]interface{}{
			"id":       hostSingle.ID,
			"hostName": hostSingle.Hostname,
			"CpuCores": hostSingle.CpuCores,
			"OS":       hostSingle.OS,
			"Status":   hostSingle.Status,
		}
		hostList = append(hostList, hostMap)
	}
	return total, hostList
}
