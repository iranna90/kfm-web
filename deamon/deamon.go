package main

import (
	"net/http"
	"log"
	"html/template"
	"kmf-frontend/dto"
	"fmt"
	"encoding/json"
	"io"
	"kmf-frontend/data"
)

var tpl *template.Template

var host = fmt.Sprintf("http://localhost:1234/kmf/dairies/khajuri")

func init() {
	tpl = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/persondetails", personOperations)
	http.HandleFunc("/registration", provideRegistrationPage)
	http.HandleFunc("/persons", handlePerson)

	log.Println("Starting server for front-end")
	http.ListenAndServe(":5678", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml",
		data.DataToDisplay{Label: "home"})
}

func provideRegistrationPage(w http.ResponseWriter, r *http.Request) {
	registration := true;
	tpl.ExecuteTemplate(w, "person-details.gohtml",
		data.DataToDisplay{IsRegistration: &registration, Label: "registration", DataHeader: "Person Registration"})
}

func personOperations(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "person-details.gohtml",
		data.DataToDisplay{Label: "persons", DataHeader: "Person Operations"})
}

func handlePerson(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

	}

	if r.Method == "GET" {

	}
	personId := r.FormValue("personId")
	url := host + "/persons/" + personId
	fmt.Println("Url is ", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, "Error while reading person", http.StatusInternalServerError)
		log.Println(err.Error())
		tpl.ExecuteTemplate(w, "person-details.gohtml", nil)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Person not found", http.StatusNotFound)
		tpl.ExecuteTemplate(w, "person-details.gohtml", nil)
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
