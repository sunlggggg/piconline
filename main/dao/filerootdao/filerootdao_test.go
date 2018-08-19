package filerootdao

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
)

func TestRootDirID(t *testing.T) {
	mysqlconfig.Init()
	fileroot, err := RootDirID(49)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(*fileroot)
	}
}
