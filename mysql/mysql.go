package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "xiaozhi:xiaozhi123@tcp(193.112.46.144:3306)/xiaozhi?charset=utf8")
	checkErr(err)
	rows, e := db.Query("select * from sys_user");
	checkErr(e)
	columns, _ := rows.Columns()
	for key, value := range columns {
		fmt.Println(key)
		fmt.Println(value)
	}
}

func checkErr(err error) {

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}