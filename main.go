package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

var bot *tgbotapi.BotAPI

// стартовое inline меню
func startMenu() tgbotapi.InlineKeyboardMarkup {
	btnSkills := tgbotapi.NewInlineKeyboardButtonData("Мои навыки", "skills") // tgbotapi.NewKeyboardButton("Привет")

	row := tgbotapi.NewInlineKeyboardRow(btnSkills)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)
	return keyboard
}

// стартовая клава
var keyboardStart = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("start"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Интересует эта неделя"),
		tgbotapi.NewKeyboardButton("Интересует определённая дата"),
	),
)

// клава для вывода информации на текущую неделю
var keyboardWeek = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Выдать инфомацию"),
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

// клава для ввода даты пользователем
var keyboardDate = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Ввести дату"),
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

// начальная функция при запуске бота
func startBot(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, меня зовут бот Боба. Хочешь узнать что я умею?")
	msg.ReplyMarkup = startMenu()
	msg.ParseMode = "Markdown"
	sendMessage(msg)
}

func addHW(update tgbotapi.Update) {
	// add homework or info in db
}

func commands(update tgbotapi.Update) { //функция которая будет реагировать на команды в чате
	command := update.Message.Command()

	switch command {
	case "start":
		startBot(update)
	case "add":
		addHW(update)
	}
}

// нажатие на ReplyKeyboard
func pressKeyboard(update tgbotapi.Update) {
	command := update.Message.Text

	switch command {
	case "start":
		startBot(update)
	case "Интересует эта неделя":
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Нажмите на кнопку, если хотите получить информацию")
		msg.ReplyMarkup = keyboardWeek
		sendMessage(msg)
	case "Интересует определённая дата":
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Нажмите на кнопку, чтобы ввести дату с домашним заданием")
		msg.ReplyMarkup = keyboardDate
		sendMessage(msg)
	}
}

// ответ на нажатие inlineButton
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

// отправка сообщений пользователю
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

	var keyWords = [](string){
		"start", "Интересует эта неделя", "Интересует определённая дата",
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			println("use callback")
			callbacks(update)
		} else if slices.Contains(keyWords, update.Message.Text) {
			pressKeyboard(update)
		} else if update.Message.IsCommand() {
			commands(update)
		} else {
			database()
			println("simple message")
		}

		if update.Message == nil {
			continue
		}
	}
}
