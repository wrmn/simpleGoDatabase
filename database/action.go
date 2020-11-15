package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
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
	db, e := sql.Open("mysql", "test:PassworD12312312?@tcp(127.0.0.1:3306)/world_x")
	ErrorCheck(e)
	return db
}

func Read() []Post {
	db := initdb()

	rows, e := db.Query(`
			SELECT
				country.code,
				country.Name as name,
				city.Name as capital
			from
				country
				left join city on city.id = country.Capital
			where
				country.name is not null
			and
				city.name is not null
	`)
	ErrorCheck(e)

	var post = Post{}
	var result []Post

	for rows.Next() {
		e = rows.Scan(&post.Code, &post.Name, &post.Capital)
		ErrorCheck(e)
		result = append(result, post)
	}
	return result
}
