package filedao

import (
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"log"
)

// 插入
func InsertFile(userID uint64, fatherID uint64, contentID uint64, createTime uint64) (int64) {
	db := mysqlconfig.Mysqldb
	res, err := db.Exec("insert into  file (userID, fatherID, isFile,  contentID, updateTime ,createTime) values (?,?,?,?,?,?);",
		userID, fatherID, true, contentID, createTime, createTime)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}

// 插入
func InsertDir(userID uint64, fatherID uint64, createTime uint64) (int64) {
	db := mysqlconfig.Mysqldb
	res, err := db.Exec("insert into  file (userID, fatherID, isFile,  contentID, updateTime ,createTime) values (?,?,?,?,?,?)",
		userID, fatherID, false, nil, createTime, createTime)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}

// 插入
func InsertRootDir(userID uint64, createTime uint64) (int64) {
	db := mysqlconfig.Mysqldb
	res, err := db.Exec("insert into  file (userID, isFile,  contentID, updateTime ,createTime) values (?,?,?,?,?)",
		userID,
		false,
		nil,
		createTime,
		createTime)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}
