package tests

import (
	"encoding/json"
	"github.com/cordova2009/go-mysql-orm/mysql"
	"github.com/cordova2009/go-mysql-orm/tests/domain"
	"strconv"
	"testing"

	"github.com/cordova2009/go-mysql-orm/utils/logger"
)

func init() {
	Db, _ := mysql.New("mysql", "root:power666@tcp(127.0.0.1:3306)/data")
	logger.Sugar.Debug(Db.Ping())
}

func TestQuery(t *testing.T) {
	//user query
	whs, err := domain.GetUser()
	if err != nil {
		return
	}
	//for _, wh := range whs {
	//	logger.Sugar.Info(wh)
	//}

	b, _ := json.Marshal(whs)
	logger.Sugar.Info(string(b))
}

func TestSave(t *testing.T) {

	data := new(domain.User)
	for i := 0; i < 10; i++ {
		data.Code = "a102asd102ad" + strconv.Itoa(i)
		data.Name = "zhangsan" + strconv.Itoa(i)
		_, err := mysql.Save(data)

		if err != nil {
			logger.Sugar.Info(err)
			return
		}

	}

}

//
//func TestUpdate(b *testing.T) {
//	//user update  add by cordova2009
//	whs, err := domain.GetUser()
//	if err != nil {
//		return
//	}
//	for _, wh := range whs {
//		data := new(domain.User)
//		code := wh
//		data.Name = wh.Name + "002"
//		//update data for mysql
//		if err = mysql.Update(data, "wh_code", code); err != nil {
//			return
//		}
//
//		logger.Sugar.Info("更新数据:%v成功", wh)
//	}
//
//}
