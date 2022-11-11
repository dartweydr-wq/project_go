package main

import (
	"html/template"
	"net/http"

	"github.com/dartweydr-wq/project_go/pkg"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html") // #TODO обработать ошибки
	tmpl.Execute(w, nil)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/contacts.html") // #TODO обработать ошибки
	tmpl.Execute(w, nil)
}

func handleReq() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//handleReq()
	pkg.ConnectionDb()
}
