package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Text    string `json:"text"`
}

var NoteStorage []Note

func main() {
	http.HandleFunc("/", HandlerHome)
	http.HandleFunc("/save", HandlerSaveNote)
	http.HandleFunc("/list_all", HandlerListAll)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func HandlerHome(resW http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resW, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	name := req.URL.Query().Get("name")
	fmt.Fprintf(resW, "Hello, %s", name)
}

func HandlerSaveNote(resW http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resW, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	n := Note{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(resW, "Bad Request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &n)
	if err != nil {
		fmt.Println(err)
		http.Error(resW, "Bad Request", http.StatusBadRequest)
		return
	}

	NoteStorage = append(NoteStorage, n)
	fmt.Fprintf(resW, "Name %s\n", n.Name)
	fmt.Fprintf(resW, "Surname %s\n", n.Surname)
	fmt.Fprintf(resW, "Note %s\n", n.Text)
}

func HandlerListAll(resW http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resW, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	Data, err := json.Marshal(NoteStorage)
	if err != nil {
		fmt.Println(err)
		http.Error(resW, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	resW.Header().Set("Content-Type", "application/json")
	resW.Write(Data)
}
