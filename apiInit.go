package main

import (
	"encoding/json"
	"fmt"
	database "local/simpleGoDatabase/database"
	"log"
	"net/http"
)

type Response struct {
	Code     int                `json:"status"`
	Response []database.Country `json:"response"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/countries", returnAllCountries)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllCountries(w http.ResponseWriter, r *http.Request) {
	Countries := database.Read()
	response := Response{
		Code:     200,
		Response: Countries,
	}
	fmt.Println("Endpoint Hit: reutrnAllCountries")
	json.NewEncoder(w).Encode(response)
}

func main() {
	handleRequests()
}
