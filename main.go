package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

var posts = []Article{}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/example_go")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	// выборка данных
	res, err := db.Query("SELECT * FROM articles ")
	defer res.Close()

	if err != nil {
		panic(err)
	}

	posts = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	tmpl.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/example_go")
		defer db.Close()

		if err != nil {
			panic(err)
		}

		// добавление данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (title, anons, full_text) VALUES('%s','%s','%s')", title, anons, full_text))
		defer insert.Close()

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleReq() {
	http.HandleFunc("/", home)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleReq()
}
