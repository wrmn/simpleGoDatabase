package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	Code             string     `json:"code"`
	Name             string     `json:"name"`
	Capital          string     `json:"capital"`
	OfficialLanguage string     `json:"officialLanguage"`
	Cities           []City     `json:"cities"`
	Languages        []Language `json:"languages"`
}

type City struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	District string `json:"dsitrict"`
}

type Language struct {
	Language   string  `json:"language"`
	IsOfficial bool    `json:"isOfficial"`
	Percentage float32 `json:"percentage"`
}

var (
	country  = Country{}
	city     = City{}
	language = Language{}
)

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

func ReadCountryLanguages(code string, db *sql.DB) []Language {
	languages := []Language{}
	rowsLanguage, e := db.Query(`
		SELECT
			language,
			isOfficial,
			percentage
		from 
			countrylanguage
		where CountryCode = ?
	`, code)
	errorCheck(e)

	for rowsLanguage.Next() {
		e = rowsLanguage.Scan(&language.Language, &language.IsOfficial, &language.Percentage)
		errorCheck(e)
		languages = append(languages, language)
	}

	return languages
}

func ReadCountryCities(code string, db *sql.DB) []City {
	cities := []City{}
	rowsCity, e := db.Query(`
		SELECT
			ID,
			Name,
			District
		FROM
			city
		WHERE
			CountryCode = ?
	`, code)
	errorCheck(e)

	for rowsCity.Next() {
		e = rowsCity.Scan(&city.Id, &city.Name, &city.District)
		errorCheck(e)
		cities = append(cities, city)
	}

	return cities

}

// read data country
func ReadCountries() []Country {
	db := initdb()
	result := []Country{}
	rowsCountry, e := db.Query(`
			SELECT
				country.Code,
				country.Name as name,
				city.Name as capital,
				countrylanguage.Language
			FROM
				country
				LEFT JOIN city ON city.ID = country.Capital
				RIGHT JOIN countrylanguage ON countrylanguage.CountryCode = country.Code
			WHERE
				country.Name IS NOT NULL
			AND
				city.Name IS NOT NULL
			AND countrylanguage.IsOfficial 
	`)
	errorCheck(e)

	for rowsCountry.Next() {
		e = rowsCountry.Scan(&country.Code, &country.Name, &country.Capital, &country.OfficialLanguage)
		errorCheck(e)
		country.Cities = ReadCountryCities(country.Code, db)
		country.Languages = ReadCountryLanguages(country.Code, db)
		result = append(result, country)
	}
	return result
}
