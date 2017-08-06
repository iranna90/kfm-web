package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"html/template"
	"kmf-frontend/dto"
	"fmt"
	"encoding/json"
	"io"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/index", handleIndex).Methods("GET")
	route.HandleFunc("/person", handlePerson).Methods("GET")
	log.Println("Starting server for front-end")
	http.ListenAndServe(":5678", route)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func handlePerson(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://ddb2455b.ngrok.io/kmf/persons/abc")
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, "Error while reading person", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	var person dto.Person
	err = decode(resp.Body, &person)

	if err != nil {
		http.Error(w, "Error while reading data from response", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	fmt.Println(person)
	tpl.ExecuteTemplate(w, "person-details.gohtml", person)
}

func decode(data io.ReadCloser, dataType interface{}) (err error) {
	decoder := json.NewDecoder(data)
	err = decoder.Decode(dataType)
	return
}
