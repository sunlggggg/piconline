package filedao

import (
	"github.com/sunlggggg/piconline/main/config/mysql"
	"strconv"
	"time"
)

func CreateDir() (int64, error) {
	db := mysql.Mysqldb
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := db.Exec("insert into file(isFile, createTime, updateTime) values (false , '" + strTime + "' , '" + strTime + "')")

	if err != nil {
		return -1, err
	} else {
		return res.LastInsertId()
	}
}
