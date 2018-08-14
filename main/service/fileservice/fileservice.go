package fileservice

import (
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"github.com/sunlggggg/piconline/main/dao/filedao"
	"github.com/sunlggggg/piconline/main/dao/filerootdao"
	"github.com/sunlggggg/piconline/main/dao/userdao"
)

func CreateRoot(username string) (int64, error) {
	// 事务测试 ...
	// 注意 在使用中应该使用tx而不是db
	user := userdao.FindByName(username)
	db := mysqlconfig.Mysqldb
	tx, _ := db.Begin()
	fileId, err := filedao.CreateDir(tx)
	if err != nil {
		tx.Rollback()
		tx.Commit()
		return -1, err
	}
	id, err := filerootdao.CreateRoot(tx, *user, fileId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	return id, err
}
