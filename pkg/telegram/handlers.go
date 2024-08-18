package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаком с такой командой :(")

	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввёл команду старт"
		_, err := b.bot.Send(msg)
		return err

	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	logrus.Infof("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	_, err := b.bot.Send(msg)
	if err != nil {
		logrus.Error(err)
	}
}
