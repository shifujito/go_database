package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var qry string = "select * from mydata where id = ?	"

type Mydatas struct {
	ID int
	Name string
	Mail string
	Age int
}

func (m *Mydatas) Str() string {
	return strconv.Itoa(m.ID) + " " + m.Name + " " + m.Mail + " " + strconv.Itoa(m.Age)
}

func input(msg string)string{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}

func search(){
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil{
		panic(er)
	}
	defer con.Close()

	for {
		s := input("id")
		if s == ""{
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil{
			panic(er)
		}
		rs, er := con.Query(qry, n)
		if er != nil{
			panic(er)
		}
		for rs.Next(){
			var md Mydatas
			er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
			if er != nil{
				panic(er)
			}
			fmt.Println(md.Str())
		}
	}
}
