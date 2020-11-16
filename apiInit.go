package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	dbCon "simpleGoDatabase/databaseConnection"
)

type ResponseCountry struct {
	Code     int             `json:"status"`
	Message  string          `json:"message"`
	Response []dbCon.Country `json:"response"`
}

type ResponseCity struct {
	Code     int          `json:"status"`
	Message  string       `json:"message"`
	Response []dbCon.City `json:"response"`
}

type ResponseLanguage struct {
	Code     int              `json:"status"`
	Message  string           `json:"message"`
	Response []dbCon.Language `json:"response"`
}

var (
	statusCode   int
	dbConnection *sql.DB
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	dbConnection = dbCon.Initdb()
	http.HandleFunc("/", homePage)
	http.HandleFunc("/countries", countries)
	http.HandleFunc("/cities", cities)
	http.HandleFunc("/languages", languages)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func status(method string) int {
	switch method {
	case "GET":
		return 200
	case "POST":
		return 201
	case "PUT":
		return 202
	case "DELETE":
		return 200
	default:
		return 404
	}
}

func errHandle(err error) string {
	fmt.Println("DB disconnect")
	fmt.Print(err)
	return "Database Disconnect, Contact your network administrator"
}

func countries(w http.ResponseWriter, r *http.Request) {
	response := ResponseCountry{}
	w.Header().Set("Content-Type", "application/json")

	err := dbCon.PingDB(dbConnection)

	if err != nil {
		response.Code = 500
		response.Message = errHandle(err)
	} else {
		statusCode = status(r.Method)

		w.WriteHeader(statusCode)

		Countries := dbCon.ReadCountries(dbConnection)

		response = ResponseCountry{
			Code:     statusCode,
			Message:  "Success",
			Response: Countries,
		}

		fmt.Println("Endpoint Hit: Countries")
	}

	json.NewEncoder(w).Encode(response)
}

func cities(w http.ResponseWriter, r *http.Request) {
	response := ResponseCity{}
	w.Header().Set("Content-Type", "application/json")

	err := dbCon.PingDB(dbConnection)

	if err != nil {
		response.Code = 500
		response.Message = errHandle(err)
	} else {
		statusCode = status(r.Method)

		w.WriteHeader(statusCode)

		Cities := dbCon.ReadCities(dbConnection)

		response = ResponseCity{
			Code:     statusCode,
			Message:  "Success",
			Response: Cities,
		}
		fmt.Println("Endpoint hit: Cities")
	}

	json.NewEncoder(w).Encode(response)
}

func languages(w http.ResponseWriter, r *http.Request) {
	response := ResponseLanguage{}
	w.Header().Set("Content-Type", "application/json")

	err := dbCon.PingDB(dbConnection)

	if err != nil {
		response.Code = 500
		response.Message = errHandle(err)
	} else {
		statusCode = status(r.Method)

		w.WriteHeader(statusCode)

		Languages := dbCon.ReadLanguages(dbConnection)

		response = ResponseLanguage{
			Code:     statusCode,
			Message:  "Success",
			Response: Languages,
		}
		fmt.Println("Endpoint hit: Languages")
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	handleRequests()
}
