package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func connectDatabase() {
	connection := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSLMODE"))

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO homework (subject, deadline) values ('Мат. анализ', '05.05.2004')")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}

func database() {
	connectDatabase()
}
