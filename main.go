package main

import (
	"database/sql"
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

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/example_go")
	defer db.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("Подключение к бд успешно")

	// добавление данных
	// insert, err := db.Query("INSERT INTO users (name, age) VALUES('Иван', 13),('Александр', 10),('Дмитрий', 22)")
	// defer insert.Close()

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Данные записались в бд")

	// выборка данных
	res, err := db.Query("SELECT name, age FROM users ")
	defer res.Close()

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
	}

}
