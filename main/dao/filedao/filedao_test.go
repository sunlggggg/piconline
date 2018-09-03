package filedao

import (
	"testing"
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"time"
)

func TestInsertRootDir(t *testing.T) {
	mysqlconfig.Init()
	InsertRootDir(49,uint64(time.Now().Unix()))
}

func TestInsertDir(t *testing.T) {
	mysqlconfig.Init()
	InsertDir(49,4, uint64(time.Now().Unix()))
}

func TestInsertFile(t *testing.T) {
	mysqlconfig.Init()
	InsertFile(49, 4,1,uint64(time.Now().Unix()))
}

