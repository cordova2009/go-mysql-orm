package mysql

import (
	"database/sql"
	"github.com/cordova2009/go-mysql-orm/utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB the mysql database.
	Db *sql.DB
)

func New(driverName, dataSourceName string) (*sql.DB, error) {
	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		logger.Sugar.Fatal(err)
		return nil, err
	}
	Db = database
	//设置数据库连接池
	//Db.SetMaxOpenConns(0)
	//Db.SetMaxIdleConns(0)
	logger.Sugar.Info("Connection Opened to mysql success")
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
	return Db, err
}
