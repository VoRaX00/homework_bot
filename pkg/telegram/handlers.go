package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		fmt.Println(message.Command())
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	logrus.Infof("[%s] %s", message.From.UserName, message.Text)

	if message.Photo == nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	} else {

		//photos := *message.Photo
		//var mediaGroup []interface{}
		//for i, photo := range photos {
		//	media := tgbotapi.NewInputMediaPhoto(photo.FileID)
		//
		//	if i == 0 {
		//		media.Caption = message.Caption
		//	}
		//
		//	mediaGroup = append(mediaGroup, media)
		//}
		//
		//msg := tgbotapi.NewMediaGroup(message.Chat.ID, mediaGroup)
		//_, err := b.bot.Send(msg)
		//
		//if err != nil {
		//	logrus.Error(err)
		//}
	}
	return nil
}

func (b *Bot) handleMediaGroup(message *tgbotapi.Message) error {

	//var mediaGroup []interface{}

	//var mediaGroup []interface{}

	//message := update.Message
	//if message != nil && message.MediaGroupID == userMessage.MediaGroupID {
	//	if len(message.Photo) > 0 {
	//		photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(message.Photo[len(message.Photo)-1].FileID))
	//
	//		if len(mediaGroup) == 0 && message.Caption != "" {
	//			photo.Caption = message.Caption
	//		}
	//		mediaGroup = append(mediaGroup, photo)
	//	}
	//}
	//
	//if len(mediaGroup) > 0 {
	//	msg := tgbotapi.NewMediaGroup(userMessage.Chat.ID, mediaGroup)
	//	if _, err := b.bot.Send(msg); err != nil {
	//		return err
	//	}
	//}
	return nil
}
