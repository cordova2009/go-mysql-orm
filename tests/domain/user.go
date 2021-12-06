package domain

import (
	"github.com/cordova2009/go-mysql-orm/mysql"
	"github.com/cordova2009/go-mysql-orm/utils/logger"
)

// User represents the information of user
type User struct {
	ID   string `json:"id"`
	Code string `json:"code"` // 编码
	Name string `json:"name"` // 名称
}

// GetUser returns the data of user
func GetUser() ([]User, error) {
	var (
		data []User
	)
	rows, _ := mysql.Query("select * from user limit 1")
	defer rows.Close()
	for rows.Next() {
		//解析rows 数据信息到结构体
		var user User
		if err := mysql.ObjMapping(rows, &user); err != nil {
			logger.Sugar.Infof("scan columns failed, error[%v]\n", err)
			return data, err
		}
		data = append(data, user)
	}
	return data, nil
}
