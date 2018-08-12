package userdao

import (
	"github.com/sunlggggg/piconline/main/entity"
	"github.com/sunlggggg/piconline/main/config/mysql"
	"fmt"
	"log"
	"database/sql"
	"strconv"
)

func Create(user entity.User) uint64 {
	db := mysql.Mysqldb
	res, err := db.Exec("insert into  user (name, email, password,  register_time) values ('" + user.Name + "', '" + user.Email + "',' " + user.Password + "', '" + strconv.FormatUint(user.RegisterTime,10) + "');")
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err !=nil {
		log.Fatal(err)
	}
	return uint64(id)
}
func FindByName(name string) *entity.User {
	db := mysql.Mysqldb
	rows, err := db.Query("select * from user where name = '" + name + "';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cloumns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var user *entity.User
	if rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(cloumns[i], ": ", value)
		}

		id, err := strconv.ParseInt(string(values[0]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		time, err := strconv.ParseInt(string(values[3]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		user = new(entity.User)
		user.Init(uint64(id), string(values[1]), string(values[4]), string(values[2]), uint64(time))

	}
	return user
}

func FindById(id uint64) *entity.User {
	db := mysql.Mysqldb
	rows, err := db.Query("select * from user where id = " + string(id) + ";")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cloumns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(cloumns[i], ": ", value)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
