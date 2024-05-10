package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

var bot *tgbotapi.BotAPI

func startMenu() tgbotapi.InlineKeyboardMarkup {
	btnSkills := tgbotapi.NewInlineKeyboardButtonData("Мои навыки", "skills") // tgbotapi.NewKeyboardButton("Привет")

	row := tgbotapi.NewInlineKeyboardRow(btnSkills)
	// Создаем строки с кнопками
	// row1 := []tgbotapi.InlineKeyboardButton{btnHi}
	// row2 := []tgbotapi.InlineKeyboardButton{btnBye}

	// Создаем клавиатуру из кнопок и строк
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)
	return keyboard
}

var keyboardStart = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("start"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Интересует эта неделя"),
		tgbotapi.NewKeyboardButton("Интересует определённая дата"),
	),
)

var keyboardWeek = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Выдать инфомацию"),
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var keyboardDate = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Ввести дату"),
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

func startBot(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, меня зовут бот Боба. Хочешь узнать что я умею?")
	msg.ReplyMarkup = startMenu()
	msg.ParseMode = "Markdowns"
	sendMessage(msg)
}

func commands(update tgbotapi.Update) { //функция которая будет реагировать на команды в чате
	command := update.Message.Command()

	switch command {
	case "start":
		startBot(update)
	}
}

func pressKeyboard(update tgbotapi.Update) {
	command := update.Message.Text

	switch command {
	case "start":
		startBot(update)
	case "Интересует эта неделя":

	}
}

func callbacks(update tgbotapi.Update) {
	callbackData := update.CallbackQuery.Data
	chatID := update.CallbackQuery.From.ID

	switch callbackData {
	case "skills":
		text := fmt.Sprintf("Привет %v", update.CallbackQuery.Message.Chat.FirstName)
		msg := tgbotapi.NewMessage(int64(chatID), text)
		msg.ReplyMarkup = keyboardStart
		sendMessage(msg)

	case "bye":
		text := fmt.Sprintf("Пока %v", update.CallbackQuery.Message.Chat.FirstName)
		msg := tgbotapi.NewMessage(int64(chatID), text)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		sendMessage(msg)
	}
}

func sendMessage(msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not loaded")
	}

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.

		if update.CallbackQuery != nil {
			callbacks(update)

		} else if update.Message.IsCommand() {
			commands(update)
		} else {
			println("simple message")
		}

		if update.Message == nil { // Ignore any non-Message updates
			continue
		}
	}
}
