package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// File Struct (Model)
type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Each file with content
type FileWithContent struct {
	File
	Size    int    `json:"size"`
	Content string `json:"content"`
}

// Init files var as a slice File Struct
var files []File

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	files = append(files, File{ID: "1", Name: "test.txt"})
	files = append(files, File{ID: "2", Name: "test2.txt"})
	files = append(files, File{ID: "3", Name: "test3.txt"})
	files = append(files, File{ID: "4", Name: "test4.txt"})
	files = append(files, File{ID: "5", Name: "test5.txt"})

	// Route Handlers or Endpoint
	r.HandleFunc("/api/files", getFiles).Methods("GET")
	r.HandleFunc("/api/files/{id}", readFile).Methods("GET")
	r.HandleFunc("/api/files", writeFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func writeFile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var newContent FileWithContent
	_ = json.NewDecoder(request.Body).Decode(&newContent)
	name, size, content := CreateFile(newContent.Name, newContent.Content)
	newContent.ID = strconv.Itoa(rand.Intn(1000000))
	newContent.Name = name
	newContent.Size = size
	newContent.Content = content

	files = append(files, File{ID: newContent.ID, Name: name})
	json.NewEncoder(writer).Encode(newContent)

}

func readFile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Get params
	params := mux.Vars(request)

	// Loop through files and find with id
	for _, item := range files {
		if item.ID == params["id"] {

			// Read content here
			var size int
			var content string

			if FuncCheckExist(item.Name) {
				size, content = ReadFile(item.Name)
			}

			// Add content to item interface
			itemWithContent := FileWithContent{
				File:    item,
				Size:    size,
				Content: content,
			}

			json.NewEncoder(writer).Encode(itemWithContent)
			return
		}
	}
	json.NewEncoder(writer).Encode(&File{})
}

func getFiles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(files)
}
