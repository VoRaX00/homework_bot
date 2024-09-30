package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()

	err := validate.RegisterValidation("customGroup", func(fl validator.FieldLevel) bool {
		group := fl.Field().String()
		fields := strings.Split(group, " ")
		if len(fields) != 2 || len(fields[0]) != 14 {
			return false
		}

		return true
	})

	if err != nil {
		err = fmt.Errorf("failed to validate text: %v", err)
		return err
	}

	err = validate.Var(message.Text, "required,customGroup")
	return err
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
	}

	fields := strings.Split(message.Text, " ")
	user := domain.NewUser(message.From.UserName, fields[0], fields[1])

	err := b.GetServices().IUserService.Create(*user)
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
	return nil
}
