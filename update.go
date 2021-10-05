package main

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func update(){
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	ids := input("update ID")
	id, _ := strconv.Atoi(ids)
	qry = "select * from mydata where id = ?"
	rw := con
}
