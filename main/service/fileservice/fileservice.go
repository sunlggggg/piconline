package fileservice

import (
	"github.com/sunlggggg/piconline/main/config/mysql"
	"github.com/sunlggggg/piconline/main/dao/filedao"
	"github.com/sunlggggg/piconline/main/dao/filerootdao"
	"github.com/sunlggggg/piconline/main/dao/userdao"
)

func CreateRoot(username string) (int64, error) {
	// TODO 事务测试 ...
	user := userdao.FindByName(username)
	db := mysql.Mysqldb
	tx, _ := db.Begin()
	fileId, err := filedao.CreateDir()
	if err != nil {
		tx.Rollback()
		tx.Commit()
		return -1, err
	}
	id, err := filerootdao.CreateRoot(*user, fileId)
	if err != nil {
		tx.Rollback()
		tx.Commit()
		return -1, err
	}
	err = tx.Commit()
	return id, err
}
