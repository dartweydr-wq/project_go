package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDb() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/example_go")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Подключение к бд успешно")

	// добавление данных
	insert, err := db.Query("INSERT INTO users (name, age) VALUES('Иван', 13),('Александр', 10),('Дмитрий', 22)")
	defer insert.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Данные записались в бд")
}
