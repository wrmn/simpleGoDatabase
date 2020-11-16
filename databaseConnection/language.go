package databaseConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ReadCountryLanguages(code string, db *sql.DB) []Language {
	languages := []Language{}
	rowsLang, e := db.Query(`
		SELECT
			language,
			isOfficial,
			percentage
		FROM 
			countrylanguage
		where CountryCode = ?
	`, code)
	errorCheck(e)

	for rowsLang.Next() {
		e = rowsLang.Scan(&language.Language, &language.IsOfficial, &language.Percentage)
		errorCheck(e)
		languages = append(languages, language)
	}

	return languages
}

func ReadLanguages(db *sql.DB) []Language {
	languages := []Language{}
	rowsLang, e := db.Query(`
		SELECT
			language,
			IsOfficial,
			percentage
		FROM
			countrylanguage
	`)
	errorCheck(e)

	for rowsLang.Next() {
		e = rowsLang.Scan(&language.Language, &language.IsOfficial, &language.Percentage)
		errorCheck(e)
		languages = append(languages, language)
	}
	return languages
}
