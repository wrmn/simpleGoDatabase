package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Name string
	Text string
}

//check error
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//check db conection
func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

//make conection
func initdb() *sql.DB {
	db, e := sql.Open("mysql", "test:PassworD12312312?@/go_simple")
	ErrorCheck(e)
	return db
}

func Read() {
	db := initdb()

	rows, e := db.Query("select * from something")
	ErrorCheck(e)

	var post = Post{}

	for rows.Next() {
		e = rows.Scan(&post.Id, &post.Name, &post.Text)
		ErrorCheck(e)
		fmt.Println(post)
	}
}
