package databaseConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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
