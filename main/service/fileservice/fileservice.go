package fileservice

import (
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"github.com/sunlggggg/piconline/main/dao/filedao"
	"github.com/sunlggggg/piconline/main/dao/filerootdao"
	"github.com/sunlggggg/piconline/main/dao/userdao"
)

func CreateRoot(userId uint64) (int64, error) {
	// 事务测试 ...
	// 注意 在使用中应该使用tx而不是db
	user := userdao.FindById(userId)
	db := mysqlconfig.Mysqldb
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

// 将文件结构记录到数据库
func InsertDir(fateuserId uint64, fatherId uint64, dirID uint64, dirName string) (uint64, error) {
	// 暂时智能操作一个文件或者一个文件夹
	// TODO
	return 0,nil
}
