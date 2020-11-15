package main

import (
	"encoding/json"
	"fmt"
	database "local/simpleGoDatabase/database"
	"log"
	"net/http"
)

type Language struct {
	Language   string `json:"name"`
	IsOfficial bool   `json:"isOfficial"`
	Percentage string `json:"percentage"`
}

type Country struct {
	Code     string     `json:"code"`
	Name     string     `json:"name"`
	Capital  string     `json:"Capital"`
	Language []Language `json:"language"`
}

var Countries database.Post
var Languages []Language

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
	fmt.Println("Endpoint Hit: reutrnAllCountries")
	json.NewEncoder(w).Encode(Countries)
}

func main() {
	Countries = database.Read()
	/* Languages = []Language{*/
	//Language{Language: "English",
	//IsOfficial: false,
	//Percentage: "23,4%",
	//},
	//Language{Language: "Bahasa",
	//IsOfficial: true,
	//Percentage: "76,6%",
	//},
	//}
	//Language2 := []Language{
	//Language{Language: "English",
	//IsOfficial: true,
	//Percentage: "100%",
	//},
	//}
	//Countries = []Country{
	//Country{Code: "IDN",
	//Name:     "Indonesia",
	//Capital:  "Jakarta",
	//Language: Languages,
	//},
	//Country{Code: "USA",
	//Name:     "United States of America",
	//Capital:  "Washington DC",
	//Language: Language2,
	//},
	/* }*/
	handleRequests()
}
