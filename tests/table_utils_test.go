package tests

import (
	"github.com/cordova2009/go-mysql-orm/utils"
	"github.com/cordova2009/go-mysql-orm/utils/logger"
	"testing"
)

func TestConvertToTableName(t *testing.T) {
	name := "test"
	logger.Sugar.Debugf("table name: %s", utils.GetTableName(name))
}
