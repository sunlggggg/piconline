package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
var Mysqldb *sql.DB
func Init() {
	// db init
	db, err := sql.Open("mysql", "root:123456@tcp(0.0.0.0:3306)/piconline?charset=utf8")
	if err != nil {
		println(err)
	}
	fmt.Println(db.Stats())
	Mysqldb = db
}
