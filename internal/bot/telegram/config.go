package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
)

func getCommandMenu() tgbotapi.SetMyCommandsConfig {
	menu := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     bot.CommandStart,
			Description: "Начать общение с ботом",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandAdd,
			Description: "Добавить новую запись",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandUpdate,
			Description: "Обновить запись",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandDelete,
			Description: "Удалить запись",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetAll,
			Description: "Всё дз",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetOnId,
			Description: "Получить дз по id",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetOnDate,
			Description: "Дз на дату",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetOnToday,
			Description: "Дз на сегодня",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetOnTomorrow,
			Description: "Дз на завтра",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetOnWeek,
			Description: "Дз на неделю",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandHelp,
			Description: "Инструкция",
		},
	)
	return menu
}
