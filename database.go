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

// func createKeyboard() {
// 	var keyboardStart = tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("start"),
// 		),
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Интересует эта неделя"),
// 			tgbotapi.NewKeyboardButton("Интересует определённая дата"),
// 		),
// 	)

// 	var keyboardWeek = tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Выдать инфомацию"),
// 			tgbotapi.NewKeyboardButton("Назад"),
// 		),
// 	)

// 	var keyboardDate = tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Ввести дату"),
// 			tgbotapi.NewKeyboardButton("Назад"),
// 		),
// 	)
// }
