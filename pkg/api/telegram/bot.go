package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"main.go/pkg/entity"
	"main.go/pkg/service/services"
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
	services   *services.Service
	userStates map[int64]string
	userData   map[int64]entity.Homework
}

func NewBot(bot *tgbotapi.BotAPI, service *services.Service) *Bot {
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
			break
		case waitingDescription:
			b.handleWaitingDescription(update.Message)
			break
		case waitingImages:
			b.handleWaitingImages(update.Message)
			break
		case waitingTags:
			b.handleWaitingTags(update.Message)
			break
		case waitingDeadline:
			b.handleWaitingDeadline(update.Message)

			id, err := b.services.Create(b.userData[userId])
			if err != nil {
				logrus.Errorf("failed to save homework: %v", err)
				_, err = b.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при добавлении"))
				if err != nil {
					logrus.Errorf("failed to send message: %v", err)
				}

				break
			}

			_, err = b.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Запись успешно сконфигурирована! ID: %d", id)))
			if err != nil {
				logrus.Errorf("failed to send message: %v", err)
				return
			}
			break
		default:
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					logrus.Errorf("[telegram] error when handling command: %s", err.Error())
				}
				break
			}

			if err := b.handleMessage(update.Message); err != nil {
				logrus.Errorf("[telegram] error when handling message: %s", err.Error())
			}
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates := b.bot.GetUpdatesChan(u)

	return updates
}
