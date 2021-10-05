package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func add_data(){
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	nm := input("name")
	ml := input("mail")
	age := input("age")
	ag, _ := strconv.Atoi(age)

	qry := "insert into mydata (name,mail,age) values(?,?,?)"
	con.Exec(qry, nm, ml, ag)
	showRecord(con)
}

func showRecord(con *sql.DB){
	qry = "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next(){
		fmt.Println(mydatafmRws(rs).Str())
	}
}

func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil{
		panic(er)
	}
	return &md
}
