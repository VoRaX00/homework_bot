package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main.go/pkg/entity"
	"time"
)

const (
	commandStart         = "start"
	commandAdd           = "add"
	commandUpdate        = "update"
	commandDelete        = "delete"
	commandHelp          = "help"
	commandGetAll        = "get_all"
	commandGetOnWeek     = "get_on_week"
	commandGetOnToday    = "get_on_today"
	commandGetOnTomorrow = "get_on_tomorrow"
	commandGetOnDate     = "get_on_date"
)

func (b *Bot) cmdStart(message *tgbotapi.Message) error {
	textStart := "Привет! Меня зовут Биба, я буду твоим помошником для получения домашек и иных новосотей!"
	msg := entity.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textStart,
	}

	err := b.SendMessage(msg, defaultChannel)
	return err
}

func (b *Bot) cmdAdd(message *tgbotapi.Message) error {
	b.switcher.ISwitcherAdd.Next()
	msg := entity.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите название домашней работы/записи",
	}

	err := b.SendMessage(msg, defaultChannel)
	return err
}

func homeworkToText(homework entity.HomeworkToGet) string {
	text := "Название: " + homework.Name + "\n" + "Описание: " + homework.Description + "\n" + "Дедлайн: " + homework.Deadline.String() + "\n"
	return text
}

func (b *Bot) cmdGetAll(message *tgbotapi.Message) error {
	homeworks, err := b.services.GetAll()

	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, defaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) cmdGetOnWeek(message *tgbotapi.Message) error {
	homeworks, err := b.services.GetByWeek()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, defaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) cmdGetOnToday(message *tgbotapi.Message) error {
	homeworks, err := b.services.GetByToday()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, defaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) cmdGetOnTomorrow(message *tgbotapi.Message) error {
	homeworks, err := b.services.GetByTomorrow()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, defaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) cmdGetOnDate(message *tgbotapi.Message) error {
	testDate := time.Date(2004, 5, 5, 0, 0, 0, 0, time.UTC)
	homeworks, err := b.services.GetByDate(testDate)
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, defaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) cmdUpdate(message *tgbotapi.Message) error {
	b.switcher.ISwitcherUpdate.Next()
	msg := entity.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите Id вашей записи",
	}

	err := b.SendMessage(msg, defaultChannel)
	return err
}

func (b *Bot) cmdDelete(message *tgbotapi.Message) error {
	err := b.services.Delete(1)
	if err != nil {
		msg := entity.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Ошибка удаления",
		}
		_ = b.SendMessage(msg, defaultChannel)
		return err
	}

	msg := entity.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Запись успешно удалена",
	}
	err = b.SendMessage(msg, defaultChannel)
	return err
}

func (b *Bot) cmdHelp(message *tgbotapi.Message) error {
	textHelp := "Инструкция пользования Бибой:"
	msg := entity.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textHelp,
	}
	err := b.SendMessage(msg, defaultChannel)
	return err
}

func (b *Bot) cmdDefault(message *tgbotapi.Message) error {
	msg := entity.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Я не знаком с такой командой :(",
	}
	err := b.SendMessage(msg, defaultChannel)
	return err
}
