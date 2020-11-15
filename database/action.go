package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Capital string `json:"capital"`
	Cities  []City `json:"cities"`
}

type City struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	District    string `json:"dsitrict"`
}

// check error
func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// check db conection
func pingDB(db *sql.DB) {
	err := db.Ping()
	errorCheck(err)
}

// make conection
func initdb() *sql.DB {
	db, e := sql.Open("mysql", "test:PassworD12312312?@tcp(127.0.0.1:3306)/world_x")
	errorCheck(e)
	return db
}

// read data country
func Read() []Country {
	db := initdb()

	rowsCountry, e := db.Query(`
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
	var city = City{}
	var cities []City

	for rowsCountry.Next() {
		e = rowsCountry.Scan(&country.Code, &country.Name, &country.Capital)
		errorCheck(e)

		city = City{}

		cities = []City{}
		rowsCity, e := db.Query(`
			SELECT
				id,
				name,
				CountryCode,
				district
			from
				city
			where
				CountryCode = ?
		`, country.Code)
		errorCheck(e)

		for rowsCity.Next() {
			e = rowsCity.Scan(&city.Id, &city.Name, &city.CountryCode, &city.District)
			errorCheck(e)
			cities = append(cities, city)
		}

		country.Cities = cities

		result = append(result, country)
	}
	return result
}
