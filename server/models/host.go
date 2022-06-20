package models

import "github.com/jinzhu/gorm"

// Host /**
// 定义主机结构体
type Host struct {
	gorm.Model
	Hostname string `json:"hostname" gorm:"commit:'主机名'"`
	OS       string `json:"os" gorm:"commit:'操作系统类型'"`
	CpuCores int    `json:"cpuCores" gorm:"commit:'CPU核心数'"`
	Status   bool   `json:"status" gorm:"commit:'主机状态'"`
}

var host Host

// GetHostList 获取主机列表
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

// CreateHost 创建主机
func CreateHost(host *Host) bool {
	if CheckHost(host.Hostname) {
		return false
	}
	db.Debug().Create(&host)
	return true
}

// CheckHost 检查主机是否已经存在
func CheckHost(hostname string) bool {
	checkHostRow := db.Debug().Select("id").Where(Host{Hostname: hostname}).First(&host)
	if checkHostRow.RowsAffected == 0 {
		return false
	}
	return true
}

func DeleteHost(hostname string) bool {
	deleteHost := Host{Hostname: hostname}
	if CheckHost(hostname) {
		db.Debug().Unscoped().Where("hostname = ?", hostname).Delete(&deleteHost)
		return true
	}
	return false
}

func EditHost(h *Host) bool {
	if CheckHost(h.Hostname) {
		db.Debug().Model(&auth).Updates(&h)
		return true
	}
	return false
}
