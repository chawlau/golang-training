package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:Root123456!@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println("Open faild")
		return
	} else {
		fmt.Println("connect succed")
	}
	_, err = db.Exec("insert into courses(cid,cname,tid)values(?, ?, ?)", "9", "ming", "6")
	if err != nil {
		fmt.Println(err)
		return
	}
}
