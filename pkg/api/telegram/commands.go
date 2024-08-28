package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/pkg/entity"
)

const (
	commandStart         = "start"
	commandAdd           = "add"
	commandUpdate        = "update"
	commandDelete        = "delete"
	commandHelp          = "help"
	commandGetAll        = "getAll"
	commandGetOnWeek     = "getOnWeek"
	commandGetOnToday    = "getOnToday"
	commandGetOnTomorrow = "getOnTomorrow"
	commandGetOnDate     = "getOnDate"
)

func (b *Bot) cmdStart(message *tgbotapi.Message) error {
	textStart := "Привет! Меня зовут Биба, я буду твоим помошником для получения домашек и иных новосотей!"
	msg := tgbotapi.NewMessage(message.Chat.ID, textStart)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) cmdAdd(message *tgbotapi.Message) error {
	b.switcher.ISwitcherAdd.Next()
	msg := tgbotapi.NewMessage(message.Chat.ID, "Напишите название домашней работы/записи")
	_, err := b.bot.Send(msg)
	return err
}

func homeworkInMessage(chatId int64, homework entity.Homework) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(chatId, "")

	msg.Text += homework.Name + "\n"
	msg.Text += homework.Description
	msg.Text += homework.Deadline.String() + "\n"
	for _, i := range homework.Tags {
		msg.Text += i + ", "
	}

	return msg, nil
}

func (b *Bot) cmdGetAll(message *tgbotapi.Message) error {
	homeworks, err := b.services.GetAll()

	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		msg, err := homeworkInMessage(message.Chat.ID, homework)
		if err != nil {
			return err
		}

		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return err
}

func (b *Bot) cmdGetOnWeek(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команда получения дз на неделю")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) cmdGetOnToday(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команда получения дз на неделю")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) cmdGetOnTomorrow(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команда получения дз на неделю")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) cmdGetOnDate(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команда получения дз на неделю")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) cmdUpdate(message *tgbotapi.Message) error {
	b.switcher.ISwitcherUpdate.Next()
	msg := tgbotapi.NewMessage(message.Chat.ID, "Напишите новое название вашего дз/записи или напишите /done")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) cmdDelete(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Команда удаления")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) cmdHelp(message *tgbotapi.Message) error {
	textHelp := "Инструкция пользования Бибой:"
	msg := tgbotapi.NewMessage(message.Chat.ID, textHelp)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) cmdDefault(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаком с такой командой :(")
	_, err := b.bot.Send(msg)
	return err
}
