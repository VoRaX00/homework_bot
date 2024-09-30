package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
	"strings"
)

type AskGroupHandler struct{}

func NewAskGroupHandler() *AskGroupHandler {
	return &AskGroupHandler{}
}

func validateGroup(message *tgbotapi.Message) error {
	group := message.Text
	fields := strings.Split(group, " ")
	if len(fields) != 2 {
		return fmt.Errorf("invalid group format")
	}
	if len(fields[0]) != 15 {
		fmt.Println(len(fields[0]))
		return fmt.Errorf("invalid group format")
	}
	return nil
}

func (h *AskGroupHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	if err := validateGroup(message); err != nil {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ! Введите ещё раз",
		}
		err = b.SendMessage(msg, bot.DefaultChannel)
		if err != nil {
			return err
		}
		return nil
	}

	fields := strings.Split(message.Text, " ")
	user, err := b.GetServices().IUserService.GetByUsername(message.From.UserName)

	if err != nil {
		user = *domain.NewUser(message.From.UserName, fields[0], fields[1])
		err = b.GetServices().IUserService.Create(user)
		if err != nil {
			msg := domain.MessageToSend{
				ChatId: message.Chat.ID,
				Text:   "Server error\n",
			}
			_ = b.SendMessage(msg, bot.DefaultChannel)
			return err
		}
		userId := message.From.ID
		b.GetSwitcher().Next(userId)

		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Группа успешно задана",
		}
		_ = b.SendMessage(msg, bot.DefaultChannel)
		return nil
	}

	err = b.GetServices().IUserService.Update(user)
	if err != nil {
		return err
	}

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Группа успешно задана",
	}
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return err
}
