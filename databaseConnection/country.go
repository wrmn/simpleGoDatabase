package databaseConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// read data country
func ReadCountries(db *sql.DB) []Country {
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
				AND city.Name IS NOT NULL
				AND countrylanguage.IsOfficial = true 
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
