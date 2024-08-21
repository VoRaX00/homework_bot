package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"main.go/pkg/service"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	services *service.Service
}

func NewBot(bot *tgbotapi.BotAPI, service *service.Service) *Bot {
	return &Bot{
		bot:      bot,
		services: service,
	}
}

func (b *Bot) Start() error {
	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				logrus.Errorf("[telegram] error when handling command: %s", err.Error())
			}
			continue
		}

		if update.Message.MediaGroupID != "" {
			if err := b.handleMediaGroup(update.Message); err != nil {
				logrus.Errorf("[telegram] error when handling media group: %s", err.Error())
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			logrus.Errorf("[telegram] error when handling message: %s", err.Error())
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates := b.bot.GetUpdatesChan(u)

	return updates
}
