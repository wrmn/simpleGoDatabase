package databaseConnection

import "database/sql"

// check error
func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// check db conection
func PingDB(db *sql.DB) error {
	err := db.Ping()
	return err
}

// make conection
func Initdb() *sql.DB {
	db, e := sql.Open("mysql", "test:PassworD12312312?@tcp(127.0.0.1:3306)/world_x")
	errorCheck(e)
	return db
}
