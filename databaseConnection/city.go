package databaseConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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
