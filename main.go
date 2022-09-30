package main

import (
	"html/template"
	"net/http"

	"github.com/dartweydr-wq/project_go/pkg"
)

type User struct {
	Name  string
	Age   uint16
	money int
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html") // #TODO обработать ошибки
	tmpl.Execute(w, nil)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html") // #TODO обработать ошибки
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
