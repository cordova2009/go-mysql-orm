package main

import (
	"github.com/cordova2009/go-mysql-orm/mysql"
	"github.com/cordova2009/go-mysql-orm/utils/logger"
)

func main() {
	Db, _ := mysql.New("mysql", "root:power666@tcp(127.0.0.1:3306)/data")
	logger.Sugar.Info(Db.Ping())
	select {}
}
