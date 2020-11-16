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

func countries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := ResponseCountry{}
	err := dbCon.PingDB(dbConnection)

	if err != nil {
		fmt.Println("DB connection disconnected")
		response = ResponseCountry{
			Code:    500,
			Message: "Database Disconnect, Contact your network administrator",
		}
	} else {
		statusCode = status(r.Method)

		w.WriteHeader(statusCode)

		Countries := dbCon.ReadCountries(dbConnection)

		response = ResponseCountry{
			Code:     statusCode,
			Response: Countries,
		}

		fmt.Println("Endpoint Hit: Countries")
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	handleRequests()
}
