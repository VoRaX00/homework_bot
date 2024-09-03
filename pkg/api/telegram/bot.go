package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"main.go/pkg/api/switcher"
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
	switcher   *switcher.Switcher
	userStates map[int64]string
	userData   map[int64]entity.Homework
}

func NewBot(bot *tgbotapi.BotAPI, service *services.Service) *Bot {
	statuses := []string{
		waitingName,
		waitingDescription,
		waitingImages,
		waitingTags,
		waitingDeadline,
	}

	return &Bot{
		bot:        bot,
		services:   service,
		switcher:   switcher.NewSwitcher(statuses, statuses),
		userData:   make(map[int64]entity.Homework),
		userStates: make(map[int64]string),
	}
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

func (b *Bot) create(message *tgbotapi.Message) {
	userId := message.From.ID
	id, err := b.services.Create(b.userData[userId])
	if err != nil {
		logrus.Errorf("failed to save homework: %v", err)
		_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Ошибка при добавлении"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
		return
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Запись успешно сконфигурирована! ID: %d", id)))
	if err != nil {
		logrus.Errorf("failed to send message: %v", err)
		return
	}
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch b.switcher.ISwitcherAdd.Current() {
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
			b.create(update.Message)
			break
		default:
			if update.Message.IsCommand() {
				if err := b.handleCommands(update.Message); err != nil {
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

const (
	defaultChannel     = 0
	channelInformation = 2
	channelBot         = 5
)

func (b *Bot) sendMediaGroup(message entity.MessageToSend, channel int) error {
	var mediaGroup []interface{}

	for i, photo := range message.Images {
		inputPhoto := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(photo))
		if i == 0 {
			inputPhoto.Caption = message.Text
		}
		mediaGroup = append(mediaGroup, inputPhoto)
	}

	mediaGroupCfg := tgbotapi.NewMediaGroup(message.ChatId, mediaGroup)
	if channel == channelBot {
		mediaGroupCfg.ReplyToMessageID = channelBot
	} else if channel == channelInformation {
		mediaGroupCfg.ReplyToMessageID = channelInformation
	}

	_, err := b.bot.SendMediaGroup(mediaGroupCfg)
	return err
}

func (b *Bot) sendText(message entity.MessageToSend, channel int) error {
	msg := tgbotapi.NewMessage(message.ChatId, "")
	msg.Text = message.Text

	if channel == channelBot {
		msg.ReplyToMessageID = channelBot
	} else if channel == channelInformation {
		msg.ReplyToMessageID = channelInformation
	}

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) SendHomework(homework entity.HomeworkToGet, chatId int64, channel int) error {
	text := homeworkToText(homework)
	msg := entity.MessageToSend{
		ChatId: chatId,
		Text:   text,
		Images: homework.Images,
	}

	err := b.SendMessage(msg, channel)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) SendMessage(message entity.MessageToSend, channel int) error {
	if len(message.Images) > 0 {
		return b.sendMediaGroup(message, channel)
	}
	return b.sendText(message, channel)
}
