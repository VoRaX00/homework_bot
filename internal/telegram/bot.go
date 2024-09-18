package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"homework_bot/internal/application/services"
	"homework_bot/internal/domain/models"
	"homework_bot/pkg/switcher"
)

const (
	waitingId          = "waitingId"
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
	userData   map[int64]models.Homework
}

func NewBot(bot *tgbotapi.BotAPI, service *services.Service) *Bot {
	statusesAdd := []string{
		waitingName,
		waitingDescription,
		waitingImages,
		waitingTags,
		waitingDeadline,
	}

	statusesUpdate := []string{waitingId}
	statusesUpdate = append(statusesUpdate, statusesAdd...)

	return &Bot{
		bot:        bot,
		services:   service,
		switcher:   switcher.NewSwitcher(statusesAdd, statusesUpdate),
		userData:   make(map[int64]models.Homework),
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

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	switch {
	case b.switcher.ISwitcherUpdate.Current() == waitingId:
		b.handleWaitingId(update.Message)
		break
	case b.switcher.ISwitcherAdd.Current() == waitingName || b.switcher.ISwitcherUpdate.Current() == waitingName:
		err := b.handleWaitingName(update.Message)
		if err != nil {
			logrus.Errorf("failed to handle waiting name: %v", err)
		}
		break
	case b.switcher.ISwitcherAdd.Current() == waitingDescription || b.switcher.ISwitcherUpdate.Current() == waitingDescription:
		err := b.handleWaitingDescription(update.Message)
		if err != nil {
			logrus.Errorf("failed to handle waiting description: %v", err)
		}
		break
	case b.switcher.ISwitcherAdd.Current() == waitingImages || b.switcher.ISwitcherUpdate.Current() == waitingImages:
		err := b.handleWaitingImages(update.Message)
		if err != nil {
			logrus.Errorf("failed to handle waiting images: %v", err)
		}
		break
	case b.switcher.ISwitcherAdd.Current() == waitingTags || b.switcher.ISwitcherUpdate.Current() == waitingTags:
		err := b.handleWaitingTags(update.Message)
		if err != nil {
			logrus.Errorf("failed to handle waiting tags: %v", err)
		}
		break
	case b.switcher.ISwitcherAdd.Current() == waitingDeadline || b.switcher.ISwitcherUpdate.Current() == waitingDeadline:
		err := b.handleWaitingDeadline(update.Message)
		if err == nil {
			b.create(update.Message)
		}
		break
	default:
		if update.Message.IsCommand() {
			if err := b.handleCommands(update.Message); err != nil {
				logrus.Errorf("[telegram] error when handling command: %s", err.Error())
			}
		}
		if err := b.handleMessage(update.Message); err != nil {
			logrus.Errorf("[telegram] error when handling message: %s", err.Error())
		}
		break
	}
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		go func(update tgbotapi.Update) {
			b.handleUpdate(update)
		}(update)
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

func (b *Bot) sendMediaGroup(message models.MessageToSend, channel int) error {
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

func (b *Bot) sendText(message models.MessageToSend, channel int) error {
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

func (b *Bot) SendHomework(homework models.HomeworkToGet, chatId int64, channel int) error {
	text := homeworkToText(homework)
	msg := models.MessageToSend{
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

func (b *Bot) SendMessage(message models.MessageToSend, channel int) error {
	if len(message.Images) > 0 {
		return b.sendMediaGroup(message, channel)
	}
	return b.sendText(message, channel)
}
