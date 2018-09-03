package filedao

import (
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"log"
	"database/sql"
)

// 插入
func InsertFile(tx *sql.Tx, userID uint64, fatherID uint64, contentID uint64, createTime uint64) (int64, error) {
	res, err := tx.Exec(`insert into  file (userID, fatherID, isFile,  contentID, updateTime ,createTime) values (?,?,?,?,?,?);`,
		userID, fatherID, true, contentID, createTime, createTime)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return res.LastInsertId()
}

// 插入
func InsertDir(userID uint64, fatherID uint64, createTime uint64) (int64, error) {
	db := mysqlconfig.Mysqldb
	res, err := db.Exec(`insert into  file (userID, fatherID, isFile,  contentID, updateTime ,createTime) values (?,?,?,?,?,?)`,
		userID, fatherID, false, nil, createTime, createTime)
	if err != nil {
		return -1, err
	} else {
		return res.LastInsertId()
	}
}

// 插入
func InsertRootDir(userID uint64, createTime uint64) (int64, error) {
	db := mysqlconfig.Mysqldb
	res, err := db.Exec(`insert into  file (userID, isFile,  contentID, updateTime ,createTime) values (?,?,?,?,?)`,
		userID,
		false,
		nil,
		createTime,
		createTime)
	if err == nil {
		return res.LastInsertId()
	} else {
		return -1, err
	}
}

func FindRoot(userId uint64) (uint64) {
	db := mysqlconfig.Mysqldb
	res := db.QueryRow(`select id from file where fatherID IS NULL  and userID = ?`, userId)
	var id uint64
	res.Scan(&id)
	return id
}
