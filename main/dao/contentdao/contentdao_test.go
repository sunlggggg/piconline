package contentdao

import (
	"testing"
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"time"
)

func TestInsert(t *testing.T) {
	mysqlconfig.Init()
	Insert("cont-01", false, uint64(time.Now().Unix()), "first file !", 49)
}
