package telegram

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	filepath "path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

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

	_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Название успешно добавлено! Теперь отправте описание к записи, или команду /done"))
	if err != nil {
		logrus.Errorf("Error sending message: %v", err)
	}
}

func (b *Bot) handleWaitingDescription(message *tgbotapi.Message) {
	userId := message.From.ID
	data := b.userData[userId]
	data.Description = message.Text
	b.userData[userId] = data
	b.userStates[userId] = waitingImages

	_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Описание успешно добавлено! Теперь отправте фотографии к записи, или команду /done"))
	if err != nil {
		logrus.Errorf("Error sending message: %v", err)
	}
}

func saveImage(bot *tgbotapi.BotAPI, fileId string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileId})
	if err != nil {
		return "", err
	}

	uniqueFileName := uuid.New().String() + filepath.Ext(file.FilePath)
	savePath := filepath.Join("../media/", uniqueFileName)

	fileURL := file.Link(bot.Token)
	response, err := http.Get(fileURL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return "", err
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
			logrus.Errorf("failed to save the file: %v", err)
		}

		data.Images = append(data.Images, path)
		b.userData[userId] = data

		_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Отправте изображение, или вызовите команду /done"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
	} else if message.Text == "/done" {
		_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Фотографии успешно загружены\nОтправте мне теги"+
			" к записи одной строкой разделяя слова запятой"))

		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
			return
		}
		b.userStates[userId] = waitingTags

	} else {
		_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "НЕВЕРНОЕ СООБЩЕНИЕ!\nНужно, то отправте изображение, или вызвать команду /done"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
	}
}

func validationTags(message *tgbotapi.Message) bool {
	for _, r := range message.Text {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func validationDate(message *tgbotapi.Message) bool {
	re := regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`)
	return re.MatchString(message.Text)
}

func (b *Bot) handleWaitingTags(message *tgbotapi.Message) {
	if !validationTags(message) {
		_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "НЕВЕРНОЕ СООБЩЕНИЕ"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
		return
	}

	userId := message.From.ID
	data := b.userData[userId]

	tags := strings.Split(message.Text, ",")
	data.Tags = tags
	b.userStates[userId] = waitingDeadline

	_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Теги успешно записаны!\nОтправте дату дедлайна записи. Формат:yyyy-mm-dd"))
	if err != nil {
		logrus.Errorf("failed to send message: %v", err)
		return
	}
}

func (b *Bot) handleWaitingDeadline(message *tgbotapi.Message) {
	if !validationDate(message) {
		_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "НЕВЕРНОЕ СООБЩЕНИЕ"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
		return
	}

	userId := message.From.ID
	data := b.userData[userId]

	layout := "2006-01-02"
	parsed, err := time.Parse(layout, message.Text)
	if err != nil {
		logrus.Errorf("failed to parse date: %v", err)
		return
	}

	data.Deadline = parsed
	b.userStates[userId] = ""

	id, err := b.services.Create(b.userData[userId])
	if err != nil {
		logrus.Errorf("failed to save homework: %v", err)
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Запись успешно сконфигурирована! ID: %d", id)))
	if err != nil {
		logrus.Errorf("failed to send message: %v", err)
		return
	}
}
