package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(0.0.0.0:3306)/piconline?charset=utf8")
	if err != nil {
		println(err)
	}
	fmt.Println(db)

	db.Exec("insert into  user (user_name, user_email, register_time) values ('sunlggggg','sunlggggg@gmail.com','12345');")

	rows, err := db.Query("select  * from user ")
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
		fmt.Println("------------------")
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	db.Exec("delete from user  where user_id > 0 && user_id < 100 ")

	db.Exec("update user set user_name = 'sun' where user_id = 22 ; ")
}
