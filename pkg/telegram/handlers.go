package telegram

import (
	"fmt"
	"io"
	"net/http"
	"os"
	filepath "path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	uuid "github.com/google/uuid"
)

const (
	commandStart  = "start"
	commandAdd    = "add"
	commandUpdate = "update"
	commandDelete = "delete"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаком с такой командой :(")

	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввёл команду старт"
		_, err := b.bot.Send(msg)

		return err
	case commandAdd:
		msg.Text = "Ты ввёл команду добавить"
		_, err := b.bot.Send(msg)

		return err
	case commandUpdate:
		msg.Text = "Ты ввёл команду обновить"
		_, err := b.bot.Send(msg)

		return err
	case commandDelete:
		msg.Text = "Ты ввёл команду удалить"
		_, err := b.bot.Send(msg)

		return err
	default:
		fmt.Println(message.Command())
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	//if message.Photo == nil {
	//	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	//	_, err := b.bot.Send(msg)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//
	//	//photos := *message.Photo
	//	//var mediaGroup []interface{}
	//	//for i, photo := range photos {
	//	//	media := tgbotapi.NewInputMediaPhoto(photo.FileID)
	//	//
	//	//	if i == 0 {
	//	//		media.Caption = message.Caption
	//	//	}
	//	//
	//	//	mediaGroup = append(mediaGroup, media)
	//	//}
	//	//
	//	//msg := tgbotapi.NewMediaGroup(message.Chat.ID, mediaGroup)
	//	//_, err := b.bot.Send(msg)
	//	//
	//	//if err != nil {
	//	//	logrus.Error(err)
	//	//}
	//}
	return nil
}

func (b *Bot) handleWaitingName(message *tgbotapi.Message) {
	userId := message.From.ID
	data := b.userData[userId]
	data.Name = message.Text
	b.userData[userId] = data
	b.userStates[userId] = waitingDescription
}

func (b *Bot) handleWaitingDescription(message *tgbotapi.Message) {
	userId := message.From.ID
	data := b.userData[userId]
	data.Description = message.Text
	b.userData[userId] = data
	b.userStates[userId] = waitingImages
}

func saveImage(bot *tgbotapi.BotAPI, fileId string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileId})
	if err != nil {
		return "", fmt.Errorf("failed to get the file: %v", err)
	}

	uniqueFileName := uuid.New().String() + filepath.Ext(file.FilePath)
	savePath := filepath.Join("../media/", uniqueFileName)

	fileURL := file.Link(bot.Token)
	response, err := http.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("failed to download the file: %v", err)
	}

	defer response.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save the file: %v", err)
	}

	return savePath, nil
}

func (b *Bot) handleWaitingImages(message *tgbotapi.Message) {
	userId := message.From.ID
	data := b.userData[userId]

	if len(message.Photo) > 0 {
		image := message.Photo[len(message.Photo)-1]
		path, err := saveImage(b.bot, image.FileID)
		if err != nil {
			fmt.Errorf("failed to save the file: %v", err)
		}
		data.Images = append(data.Images, path)
	}
	b.userStates[userId] = waitingImages
}
