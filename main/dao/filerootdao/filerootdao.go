package filerootdao

import (
	"github.com/sunlggggg/piconline/main/entity"
	"strconv"
	"time"
	"database/sql"
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
)


func CreateRoot(tx *sql.Tx, user entity.User, fileId int64) (int64, error) {

	res, err := tx.Exec("insert into fileroot (fileroot.userId,fileroot.createTime, fileroot.fileId) values  ('" + strconv.FormatUint(user.Id, 10) + "', '" + strconv.FormatInt(time.Now().Unix(), 10) + "', '" + strconv.FormatInt(fileId, 10) + "');")
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// 用户根目录是否存在
func RootDirID(userID uint64) (*entity.Fileroot, error) {
	db := mysqlconfig.Mysqldb
	rows := db.QueryRow("select id, userId, fileId, createTime  from fileroot where userId = '" + strconv.FormatUint(userID, 10) + "';")
	if rows != nil {
		fileroot := new(entity.Fileroot)
		err := rows.Scan(&fileroot.Id, &fileroot.UserId, &fileroot.FileId, &fileroot.CreateTime)
		if err != nil {
			return nil, err
		} else {
			return fileroot, nil
		}
	}
	return nil, nil
}

