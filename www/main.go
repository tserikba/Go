package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Data setting
	insert, err := db.Query("INSERT INTO Users (name, age) VALUES('Alexander', 25)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Println("Connected to mysql")
}
