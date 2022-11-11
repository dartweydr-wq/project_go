package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "index", nil)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/contacts.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.Execute(w, nil)
}

func handleReq() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleReq()
}
