package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type fileData struct {
	Name    string `json:"name"`
	Size    int    `json:"size"`
	Content string `json:"content"`
}

type fileRequest struct {
	Name string `json:"name"`
}

func readHandler(w http.ResponseWriter, r *http.Request) {

	name := fileRequest{}
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("error while reading r.Body: ", err)
	}
	err = json.Unmarshal(jsn, &name)
	if err != nil {
		log.Fatal("decoding error: ", err)
	}
	log.Printf("received: %v\n", name)

	if FuncCheckExist(name.Name) {
		fileSize, fileContent := ReadFile(name.Name)
		textFile := fileData{Size: fileSize, Content: fileContent}
		textJson, err := json.MarshalIndent(textFile, "", "  ")
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(textJson)
	} else {
		fmt.Printf("\nFile not found\n")
	}

}

func FuncCheckExist(namaFile string) bool {
	info, err := os.Stat("storages/" + namaFile)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadFile(fileName string) (int, string) {

	data, err := ioutil.ReadFile("storages/" + fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return len(data), string(data)
}

func CreateFile(fileName string, content string) (string, int, string) {

	if !strings.Contains(fileName, ".txt") {
		fileName += ".txt"
	}

	file, err := os.Create("storages/" + fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString(content)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	return fileName, len, content

}

func server() {
	http.HandleFunc("/", readHandler)
	http.ListenAndServe(":8080", nil)
}

func client() {
	var namaFile string
	fmt.Println("Type your file's name")
	fmt.Scanln(&namaFile)
	reqJson, err := json.Marshal(fileRequest{Name: namaFile})
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

//func main()  {
//	// fmt.Println("Hello World")
//
//	// textFile := fileData{Name: "Nama File", Size: 14, Data: "Some content of the text files"}
//
//	// fmt.Printf("%+v\n", textFile)
//
//	// byteArray, err := json.MarshalIndent(textFile, "", "  ")
//	// if err != nil {
//	// 	fmt.Println(err)
//	// }
//
//	// fmt.Println(string(byteArray))
//
//		go server()
//		client()
//
//}
