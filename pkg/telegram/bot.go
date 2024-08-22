package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"main.go/entity"
	"main.go/pkg/service"
)

const (
	waitingName        = "waitingName"
	waitingDescription = "waitingDescription"
	waitingImages      = "waitingImages"
	waitingTags        = "waitingTags"
	waitingDeadline    = "waitingDeadline"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	services   *service.Service
	userStates map[int64]string
	userData   map[int64]entity.Homework
}

func NewBot(bot *tgbotapi.BotAPI, service *service.Service) *Bot {
	return &Bot{
		bot:        bot,
		services:   service,
		userData:   make(map[int64]entity.Homework),
		userStates: make(map[int64]string),
	}
}

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
	)

	return menu
}

func (b *Bot) Start() error {
	_, err := b.bot.Request(getCommandMenu())
	if err != nil {
		return err
	}
	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		userId := update.Message.From.ID

		switch b.userStates[userId] {
		case waitingName:
			b.handleWaitingName(update.Message)
		case waitingDescription:
			b.handleWaitingDescription(update.Message)
		case waitingImages:
			b.handleWaitingImages(update.Message)
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				logrus.Errorf("[telegram] error when handling command: %s", err.Error())
			}
			continue
		}

		//if update.Message.MediaGroupID != "" {
		//	if err := b.handleMediaGroup(update.Message); err != nil {
		//		logrus.Errorf("[telegram] error when handling media group: %s", err.Error())
		//	}
		//	continue
		//}

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
