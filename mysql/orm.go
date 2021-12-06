package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/cordova2009/go-mysql-orm/utils"
	"github.com/cordova2009/go-mysql-orm/utils/logger"
)

// Define common vars
var (
	ErrArgs = errors.New("args error may be empty")
)

// Query query data
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = Db.Query(query, args...)

	//close the rows collections
	//defer rows.Close()
	if err != nil {
		fmt.Printf("query data from MYSQL failed, error[%v]\n", err)
		return rows, err
	}

	return rows, nil
}

// Save save data
func Save(objects interface{}, isClean ...bool) (int64, error) {

	var (
		cnt       int64
		isReplace bool
		phs       []string //占位符
		fields    []string
		values    []interface{}
	)
	if len(isClean) == 0 {
		isReplace = false
	} else {
		isReplace = isClean[0]
	}
	tableName := utils.GetTableName(objects)
	logger.Sugar.Debugf("Table name: %s", tableName)

	start := time.Now()
	if isReplace {
		Delete(objects)
	}

	AnalysisFeilds(objects, &fields, &values, &phs)

	field := strings.Join(fields, ",")
	ph := strings.Join(phs, ",")

	_, err := Db.Exec("Insert into "+tableName+" ("+field+") values ("+ph+")", values...)
	if err != nil {
		return cnt, err
	}
	logger.Sugar.Debugf("save data time: %.2f s", time.Since(start).Seconds())
	return cnt, nil
}

// Update update data
func Update(objects interface{}, idField string, idValue string) error {

	tableName := utils.GetTableName(objects)
	logger.Sugar.Debugf("Table name: %s", tableName)

	var (
		fields []string
		values []interface{}
		err    error
	)

	AnalysisFeilds(objects, &fields, &values)

	if len(fields) > 0 {
		field := strings.Join(fields, " = ?, ")
		SQL := fmt.Sprintf("Update %s set %s = ? where %s = '%s'", tableName, field, idField, idValue)

		_, err = Db.Exec(SQL, values...)
	}

	if err != nil {
		return err
	}
	return nil
}

// Delete clean data
func Delete(objects interface{}, id ...string) (int64, error) {

	var (
		cnt int64
		sql string
		phs []string //占位符
	)

	tableName := utils.GetTableName(objects)
	logger.Sugar.Debugf("Table name: %s", tableName)

	sql = fmt.Sprintf("delete from %s", tableName)
	tx, err := Db.Begin()
	if err != nil {
		return cnt, err
	}
	if len(id) == 0 {
		sql = fmt.Sprintf("delete from %s", tableName)
		_, err := tx.Exec(sql)
		if err != nil {
			return cnt, err
		}
	} else {
		for j := 0; j < len(id); j++ {
			phs = append(phs, "?")
		}
		ph := strings.Join(phs, ",")
		sql = "delete from " + tableName + " where id in(" + ph + ")"
		_, err := tx.Exec(sql, id)
		if err != nil {
			return cnt, err
		}
	}
	if err = tx.Commit(); err != nil {
		return cnt, err
	}
	return cnt, nil
}

//对象解析映射
func ObjMapping(rows *sql.Rows, structPtr interface{}) error {
	v := reflect.ValueOf(structPtr)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("dest should be a struct's pointer")
	}
	e := v.Elem()
	t := e.Type()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	var dest_scans = make([]interface{}, len(cols))
	for i, c := range cols {
		for j := 0; j < t.NumField(); j++ {
			if t.Field(j).Tag.Get("json") == c {
				dest_scans[i] = e.Field(j).Addr().Interface()
			}
		}
	}
	rows.Scan(dest_scans...)
	return nil

}

func AnalysisFeilds(objects interface{}, fields *[]string, values *[]interface{}, pis ...*[]string) error {
	ind := reflect.Indirect(reflect.ValueOf(objects))
	var (
		phs []string
	)

	flen := ind.NumField()

	for j := 0; j < flen; j++ {
		t := reflect.TypeOf(ind.Interface())
		value := reflect.ValueOf(ind.Interface())
		var name string
		if t.Field(j).Tag.Get("db") == "" {
			name = t.Field(j).Tag.Get("json")
		} else {
			name = t.Field(j).Tag.Get("db")
		}
		//ignore empty value
		if value.Field(j).IsZero() {
			continue
		}

		*fields, phs = append(*fields, name), append(phs, "?")

		v := value.Field(j).Interface()
		//The assertion determines whether it is an array type
		val, ok := v.([]string)
		if ok {
			*values = append(*values, strings.Join(val, ","))
		} else {
			//反射判断 如果为struct类型 则将字段转化为json
			if reflect.TypeOf(v).Kind() == reflect.Struct {
				str, err := json.Marshal(v)
				if err != nil {
					logger.Sugar.Fatal(err)
				}
				*values = append(*values, string(str))
				continue
			}
			*values = append(*values, v)
		}

		logger.Sugar.Debugf("Object %v: {%s:%#v(%s)}", t, name, v, reflect.TypeOf(value.Field(j).Interface()))
	}
	if len(pis) > 0 {
		*pis[0] = phs
	}
	return nil

}
