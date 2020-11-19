package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Struct for update or write new files
type UpdateFile struct {
	FileWithContent
	Replace int `json:"replace"`
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
	files = append(files, File{ID: "6", Name: "test6.txt"})
	files = append(files, File{ID: "7", Name: "test7.txt"})

	// Route Handlers or Endpoint
	r.HandleFunc("/api/files", getFiles).Methods("GET")
	r.HandleFunc("/api/files/{id}", readFile).Methods("GET")
	r.HandleFunc("/api/files", writeFile).Methods("PUT")
	r.HandleFunc("/api/checkId", check).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func check(writer http.ResponseWriter, request *http.Request) {
	var totalId string
	totalId = strconv.Itoa(len(files) + 1)

	json.NewEncoder(writer).Encode(totalId)
}

func writeFile(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var newContent UpdateFile
	_ = json.NewDecoder(request.Body).Decode(&newContent)

	//Check if file's name already exist
	fileName := newContent.Name
	if !strings.Contains(fileName, ".txt") {
		fileName += ".txt"
	}

	found := false
	var foundId string
	for i := 0; i < len(files) && found == false; i++ {
		if fileName == files[i].Name {
			found = true
			foundId = files[i].ID
		}
	}

	if found == false {
		name, size, content := CreateFile(fileName, newContent.Content)
		//newContent.ID = strconv.Itoa(rand.Intn(1000000))
		newContent.ID = strconv.Itoa(len(files) + 1)
		newContent.Name = name
		newContent.Size = size
		newContent.Content = content
		files = append(files, File{ID: newContent.ID, Name: name})
	} else {
		if newContent.Replace == 0 {
			var size int
			var content string
			size, content = ReadFile(fileName)

			name, size, content := CreateFile(fileName, content+"\n"+newContent.Content)
			newContent.ID = foundId
			newContent.Name = name
			newContent.Size = size
			newContent.Content = content

		} else {
			name, size, content := CreateFile(fileName, newContent.Content)
			newContent.ID = foundId
			newContent.Name = name
			newContent.Size = size
			newContent.Content = content

		}
	}

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
