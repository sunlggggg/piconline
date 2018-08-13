package filerootdao

import (
	"github.com/sunlggggg/piconline/main/entity"
	"strconv"
	"time"
	"database/sql"
)

type TxError struct {
}

func (t *TxError) Error() string {
	return "test tx"
}

func CreateRoot(tx *sql.Tx, user entity.User, fileId int64) (int64, error) {

	if true {
		return -1, new(TxError)
	}
	res, err := tx.Exec("insert into fileroot (fileroot.userId,fileroot.createTime, fileroot.fileId) values  ('" + strconv.FormatUint(user.Id, 10) + "', '" + strconv.FormatInt(time.Now().Unix(), 10) + "', '" + strconv.FormatInt(fileId, 10) + "');")
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
