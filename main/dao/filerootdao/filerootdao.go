package filerootdao

import (
	"github.com/sunlggggg/piconline/main/entity"
	"github.com/sunlggggg/piconline/main/config/mysql"
	"strconv"
	"time"
)

//自定义错误类型
type TxError struct {
	error //实现error接口
}

func (t *TxError) Error() string {
	return "test tx"
}

func CreateRoot(user entity.User, fileId int64) (int64, error) {
	db := mysql.Mysqldb
	res, err := db.Exec("insert into fileroot (fileroot.userId,fileroot.createTime, fileroot.fileId) values  ('" + strconv.FormatUint(user.Id, 10) + "', '" + strconv.FormatInt(time.Now().Unix(), 10) + "', '" + strconv.FormatInt(fileId, 10) + "');")
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
