package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getCommandMenu() tgbotapi.SetMyCommandsConfig {
	menu := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     commandStart,
			Description: "Начать общение с ботом",
		},
		tgbotapi.BotCommand{
			Command:     commandAdd,
			Description: "Добавить новую запись",
		},
		tgbotapi.BotCommand{
			Command:     commandUpdate,
			Description: "Обновить запись",
		},
		tgbotapi.BotCommand{
			Command:     commandDelete,
			Description: "Удалить запись",
		},
		tgbotapi.BotCommand{
			Command:     commandGetAll,
			Description: "Всё дз",
		},
		tgbotapi.BotCommand{
			Command:     commandGetOnDate,
			Description: "Дз на дату",
		},
		tgbotapi.BotCommand{
			Command:     commandGetOnToday,
			Description: "Дз на сегодня",
		},
		tgbotapi.BotCommand{
			Command:     commandGetOnTomorrow,
			Description: "Дз на завтра",
		},
		tgbotapi.BotCommand{
			Command:     commandGetOnWeek,
			Description: "Дз на неделю",
		},
		tgbotapi.BotCommand{
			Command:     commandHelp,
			Description: "Инструкция",
		},
	)
	return menu
}
