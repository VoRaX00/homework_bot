package telegram

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func isAdmin(chatId int64) bool {
	adminId, err := strconv.Atoi(os.Getenv("ADMIN4"))
	if err != nil {
		return false
	}

	return int64(adminId) == chatId
}

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	if isAdmin(message.Chat.ID) {
		switch message.Command() {
		case commandAdd:
			err := b.cmdAdd(message)
			return err
		case commandUpdate:
			err := b.cmdUpdate(message)
			return err
		case commandDelete:
			err := b.cmdDelete(message)
			return err
		}
	}

	switch message.Command() {
	case commandStart:
		err := b.cmdStart(message)
		return err
	case commandHelp:
		err := b.cmdHelp(message)
		return err
	case commandGetAll:
		err := b.cmdGetAll(message)
		return err
	case commandGetOnWeek:
		err := b.cmdGetOnWeek(message)
		return err
	case commandGetOnToday:
		err := b.cmdGetOnToday(message)
		return err
	case commandGetOnTomorrow:
		err := b.cmdGetOnTomorrow(message)
		return err
	case commandGetOnDate:
		err := b.cmdGetOnDate(message)
		return err
	default:
		err := b.cmdDefault(message)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handleWaitingName(message *tgbotapi.Message) {
	userId := message.From.ID
	data := b.userData[userId]
	data.Name = message.Text
	b.userData[userId] = data
	b.switcher.ISwitcherAdd.Next()

	err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "Название успешно добавлено! Теперь отправте описание к записи, или команду /done"))
	if err != nil {
		logrus.Errorf("Error sending message: %v", err)
	}
}

func (b *Bot) handleWaitingDescription(message *tgbotapi.Message) {
	userId := message.From.ID
	data := b.userData[userId]
	data.Description = message.Text
	b.userData[userId] = data
	b.switcher.ISwitcherAdd.Next()

	err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "Описание успешно добавлено! Теперь отправте фотографии к записи, или команду /done"))
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
	savePath := filepath.Join("/home/nikita/go/src/homework_bot/media", uniqueFileName)

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

		err = b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "Отправте изображение, или вызовите команду /done"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
	} else if message.Text == "/done" {
		err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "Фотографии успешно загружены\nОтправте мне теги"+
			" к записи одной строкой разделяя слова запятой"))

		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
			return
		}
		b.switcher.ISwitcherAdd.Next()
	} else {
		err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "НЕВЕРНОЕ СООБЩЕНИЕ!\nНужно, то отправте изображение, или вызвать команду /done"))
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
		err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "НЕВЕРНОЕ СООБЩЕНИЕ"))
		if err != nil {
			logrus.Errorf("failed to send message: %v", err)
		}
		return
	}

	userId := message.From.ID
	data := b.userData[userId]

	tags := strings.Split(message.Text, ",")
	data.Tags = tags

	b.userData[userId] = data
	b.switcher.ISwitcherAdd.Next()
	err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "Теги успешно записаны!\nОтправте дату дедлайна записи. Формат:yyyy-mm-dd"))
	if err != nil {
		logrus.Errorf("failed to send message: %v", err)
		return
	}
}

func (b *Bot) handleWaitingDeadline(message *tgbotapi.Message) {
	if !validationDate(message) {
		err := b.SendMessage(tgbotapi.NewMessage(message.Chat.ID, "НЕВЕРНОЕ СООБЩЕНИЕ"))
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

	b.userData[userId] = data
	b.switcher.ISwitcherAdd.Next()
}