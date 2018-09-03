package contentdao

import (
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"log"
)

func Insert(name string, isPublic bool, uploadTime uint64, desc string, userId uint64) (int64) {
	db := mysqlconfig.Mysqldb
	res, err := db.Exec("insert into  content (name, is_public, uploadTime, description , user_id) values (?,?,?,?,?)",
		name, isPublic, uploadTime, desc, userId)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return id
}
