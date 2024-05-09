// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// 	"github.com/joho/godotenv"
// )

// var bot *tgbotapi.BotAPI

// func startMenu() tgbotapi.InlineKeyboardMarkup {
// 	btnHi := tgbotapi.NewInlineKeyboardButtonData("Привет", "hi") // tgbotapi.NewKeyboardButton("Привет")
// 	btnBye := tgbotapi.NewInlineKeyboardButtonData("Пока", "bye") // tgbotapi.NewKeyboardButton("Пока")

// 	row := tgbotapi.NewInlineKeyboardRow(btnHi, btnBye)
// 	// Создаем строки с кнопками
// 	// row1 := []tgbotapi.InlineKeyboardButton{btnHi}
// 	// row2 := []tgbotapi.InlineKeyboardButton{btnBye}

// 	// Создаем клавиатуру из кнопок и строк
// 	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)
// 	return keyboard
// }

// func commands(update tgbotapi.Update) { //функция которая будет реагировать на команды в чате
// 	command := update.Message.Command()

// 	switch command {
// 	case "start":
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я бот и я помогу искать дз проще!!!")
// 		msg.ReplyMarkup = startMenu()
// 		sendMessage(msg)
// 	}
// }

// func callbacks(update tgbotapi.Update) {
// 	callbackData := update.CallbackQuery.Data
// 	chatID := update.CallbackQuery.Message.Chat.ID

// 	switch callbackData {
// 	case "hi":
// 		text := fmt.Sprintf("Привет %v", update.CallbackQuery.Message.Chat.FirstName)
// 		msg := tgbotapi.NewMessage(chatID, text)
// 		bot.Send(msg)

// 	case "bye":
// 		text := fmt.Sprintf("Пока %v", update.CallbackQuery.Message.Chat.FirstName)
// 		msg := tgbotapi.NewMessage(chatID, text)
// 		bot.Send(msg)
// 	}
// }

// func sendMessage(msg tgbotapi.Chattable) {
// 	if _, err := bot.Send(msg); err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

// func main() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal(".env not loaded")
// 	}

// 	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	bot.Debug = true

// 	//log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 10

// 	updates, err := bot.GetUpdatesChan(u)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	// Loop through each update.
// 	for update := range updates {
// 		// Check if we've gotten a message update.

// 		if update.Message.IsCommand() {
// 			commands(update)
// 		} else if update.CallbackQuery != nil {
// 			callbacks(update)

// 			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

// 			// if _, err := bot.Request(callback); err != nil {
// 			// 	panic(err)
// 			// }

// 			// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
// 			// if _, err := bot.Send(msg); err != nil {
// 			// 	panic(err)
// 			// }
// 		} else {
// 			println("simple message")
// 		}

// 		if update.Message == nil { // Ignore any non-Message updates
// 			continue
// 		}
// 		// } else { //если наше сообщение не команда, то мы делаем эхо
// 		// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// 		// 	bot.Send(msg)
// 		// }
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/crocone/tg-bot"
	"github.com/joho/godotenv"
)

var bot *tgbotapi.BotAPI

type button struct {
	name string
	data string
}

func startMenu() tgbotapi.InlineKeyboardMarkup {
	states := []button{
		{
			name: "Привет",
			data: "hi",
		},
		{
			name: "Пока",
			data: "buy",
		},
	}

	buttons := make([][]tgbotapi.InlineKeyboardButton, len(states))
	for index, state := range states {
		buttons[index] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(state.name, state.data))
	}

	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not loaded")
	}

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_API_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to initialize Telegram bot API: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Failed to start listening for updates %v", err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			callbacks(update)
		} else if update.Message.IsCommand() {
			commands(update)
		} else {
			// simply message
		}
	}
}

func callbacks(update tgbotapi.Update) {
	data := update.CallbackQuery.Data
	chatId := update.CallbackQuery.From.ID
	firstName := update.CallbackQuery.From.FirstName
	lastName := update.CallbackQuery.From.LastName
	var text string
	switch data {
	case "hi":
		text = fmt.Sprintf("Привет %v %v", firstName, lastName)
	case "buy":
		text = fmt.Sprintf("Пока %v %v", firstName, lastName)
	default:
		text = "Неизвестная команда"
	}
	msg := tgbotapi.NewMessage(chatId, text)
	sendMessage(msg)
}

func commands(update tgbotapi.Update) {
	command := update.Message.Command()
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие")
		msg.ReplyMarkup = startMenu()
		msg.ParseMode = "Markdown"
		sendMessage(msg)
	}
}

func sendMessage(msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		log.Panicf("Send message error: %v", err)
	}
}
