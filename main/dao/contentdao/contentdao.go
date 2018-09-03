package contentdao

import (
	"log"
	"database/sql"
)

func Insert(tx *sql.Tx, name string, isPublic bool, uploadTime uint64, desc string, userId uint64) (int64, error) {

	res, err := tx.Exec(`insert into  content (name, is_public, uploadTime, description , user_id) values (?,?,?,?,?)`,
		name, isPublic, uploadTime, desc, userId)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Println(err)
		return -1, err
	} else {
		return res.LastInsertId()
	}
}
