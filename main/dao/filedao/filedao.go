package filedao

import (
	"strconv"
	"time"
	"database/sql"
)

func CreateDir(tx *sql.Tx) (int64, error) {
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := tx.Exec("insert into file(isFile, createTime, updateTime) values (false , '" + strTime + "' , '" + strTime + "')")

	if err != nil {
		return -1, err
	} else {
		return res.LastInsertId()
	}
}

// 插入一个文件或者文件夹
func InsertFile(isFile bool, contentId uint64, time uint64) (int64, error) {
	// TODO
	return 0, nil
}
