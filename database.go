package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"main.go/homework"
)

func connectDatabase() (*sql.DB, error) {
	connection := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSLMODE"))

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	return db, nil
}

// функция определяющая является ли запрос sql инъекцией
func isSqlInjection(query string) bool {
	return false
}

// функция которая возвращает срез Homework
func getRows(rows *sql.Rows) []homework.Homework {
	homeworks := [](homework.Homework){}

	for rows.Next() {
		hw := homework.Homework{}
		var subject string
		var content string
		var deadline time.Time
		err := rows.Scan(&subject, &content, &deadline)

		if err != nil {
			panic(err)
		}

		hw.SetSubject(subject)
		hw.SetContent(content)
		hw.SetDeadline(deadline)
		homeworks = append(homeworks, hw)
	}

	return homeworks
}

// функция которая возвращает все домашние задания по выбранному предмету
func SelectSubjectHomework(subject string, db *sql.DB) ([]homework.Homework, error) {
	query := fmt.Sprintf("SELECT * FROM homework WHERE subject=%v", subject)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	return getRows(rows), nil
}

// функция которая возращает абсолютно все домашние задания
func SelectAllHomework(db *sql.DB) ([]homework.Homework, error) {
	rows, err := db.Query("SELECT * FROM homework")

	if err != nil {
		panic(err)
	}

	return getRows(rows), nil
}
