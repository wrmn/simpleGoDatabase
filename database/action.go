package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
}

//check error
func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//check db conection
func pingDB(db *sql.DB) {
	err := db.Ping()
	errorCheck(err)
}

//make conection
func initdb() *sql.DB {
	db, e := sql.Open("mysql", "test:PassworD12312312?@tcp(127.0.0.1:3306)/world_x")
	errorCheck(e)
	return db
}

//
func Read() []Country {
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
	errorCheck(e)

	var country = Country{}
	var result []Country

	for rows.Next() {
		e = rows.Scan(&country.Code, &country.Name, &country.Capital)
		errorCheck(e)
		result = append(result, country)
	}
	return result
}
