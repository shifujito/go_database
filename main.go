package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Mydata struct {
	ID int
	Name string
	Mail string
	Age int
}

func (m *Mydata) Str() string {
	return strconv.Itoa(m.ID) + " " + m.Name + " " + m.Mail + " " + strconv.Itoa(m.Age)
}

func main(){
	connect, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil{
		panic(er)
	}
	defer connect.Close()

	query := "select * from mydata"
	response, er := connect.Query(query)
	if er != nil{
		panic(er)
	}
	for response.Next(){
		var md Mydata
		er := response.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
}
