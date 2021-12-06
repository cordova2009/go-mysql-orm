package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// GetTableName convert the type name of a object to a table name. lower case
func GetTableName(obj interface{}) string {

	tab := fmt.Sprintf("%s", reflect.TypeOf(obj))
	index := strings.Index(tab, ".") + 1
	return fmt.Sprintf("%s", strings.ToLower(tab[index:]))
}
