package main

import (
	"encoding/json"
	"fmt"
	"local/simpleGoDatabase/database"
	"log"
	"net/http"
)

type ResponseCountry struct {
	Code     int                `json:"status"`
	Response []database.Country `json:"response"`
}

var (
	statusCode int
	response   []database.Country
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
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

	statusCode = status(r.Method)

	w.WriteHeader(statusCode)

	Countries := database.ReadCountries()

	response := ResponseCountry{
		Code:     statusCode,
		Response: Countries,
	}

	fmt.Println("Endpoint Hit: Countries")

	json.NewEncoder(w).Encode(response)
}

func main() {
	handleRequests()
}
