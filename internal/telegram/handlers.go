package telegram

import (
	validator "github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"homework_bot/internal/domain/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func isAdmin(chatId int64) bool {
	adminsString := os.Getenv("ADMIN4")
	adminId := strings.Split(adminsString, ",")
	for _, item := range adminId {
		id, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			logrus.Errorf("Error with convert string to int, %s", err.Error())
		}
		if id == chatId {
			return true
		}
	}

	return false
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
	case commandGetOnId:
		err := b.cmdGetOnId(message)
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

func (b *Bot) handleWaitingName(message *tgbotapi.Message) error {
	userId := message.From.ID
	data := b.userData[userId]
	data.Name = message.Text
	b.userData[userId] = data
	b.switcher.Next()

	msg := models.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Название успешно добавлено! Теперь отправте описание к записи, или команду /done",
	}

	err := b.sendMessage(msg, defaultChannel)
	return err
}

func (b *Bot) handleWaitingDescription(message *tgbotapi.Message) error {
	userId := message.From.ID
	data := b.userData[userId]
	data.Description = message.Text
	b.userData[userId] = data
	b.switcher.Next()

	msg := models.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Описание успешно добавлено! Теперь отправте фотографии к записи, или команду /done",
	}

	err := b.sendMessage(msg, defaultChannel)
	return err
}

func saveImage(bot *tgbotapi.BotAPI, fileId string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileId})
	if err != nil {
		return "", err
	}

	uniqueFileName := uuid.New().String() + filepath.Ext(file.FilePath)
	savePath := filepath.Join("/home/nikita/GolandProjects/homework_bot/media", uniqueFileName)

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

func (b *Bot) handleWaitingImages(message *tgbotapi.Message) error {
	userId := message.From.ID
	data := b.userData[userId]

	if len(message.Photo) > 0 {
		image := message.Photo[len(message.Photo)-1]
		path, err := saveImage(b.bot, image.FileID)
		if err != nil {
			return err
		}

		data.Images = append(data.Images, path)
		b.userData[userId] = data

		msg := models.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Отправте изображение, или вызовите команду /done",
		}

		err = b.sendMessage(msg, defaultChannel)
		if err != nil {
			return err
		}
	} else if message.Text == "/done" {
		msg := models.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Фотографии успешно загружены\nОтправте мне теги к записи одной строкой разделяя слова запятой",
		}

		err := b.sendMessage(msg, defaultChannel)
		if err != nil {
			return err
		}
		b.switcher.Next()
	} else {
		msg := models.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ!\nНужно, то отправте изображение, или вызвать команду /done",
		}
		err := b.sendMessage(msg, defaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

const tagsFormat = `^[a-zA-Z0-9]+(,[a-zA-Z0-9]+)*$`

func validateTags(fl validator.FieldLevel) bool {
	tags := fl.Field().String()
	matched, _ := regexp.MatchString(tagsFormat, tags)
	return matched
}

type Tags struct {
	tags string `validator:"tags"`
}

func (b *Bot) handleWaitingTags(message *tgbotapi.Message) error {
	validate := validator.New()
	err := validate.RegisterValidation("tags", validateTags)
	if err != nil {
		return err
	}

	tags := Tags{
		tags: message.Text,
	}

	if err = validate.Struct(tags); err != nil {
		msg := models.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ",
		}
		err := b.sendMessage(msg, defaultChannel)
		if err != nil {
			return err
		}

	}

	userId := message.From.ID
	data := b.userData[userId]

	tagsString := strings.Split(message.Text, ",")
	data.Tags = tagsString

	b.userData[userId] = data
	b.switcher.Next()

	msg := models.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Теги успешно записаны!\nОтправте дату дедлайна записи. Формат:yyyy-mm-dd",
	}

	err = b.sendMessage(msg, defaultChannel)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleWaitingDeadline(message *tgbotapi.Message) error {
	validate := validator.New()

	err := validate.RegisterValidation("customDate", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		_, err := time.Parse("2006-01-02", dateStr)
		return err == nil
	})
	if err != nil {
		return err
	}

	err = validate.Var(message.Text, "required,customDate")
	if err != nil {
		msg := models.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ\nВведите ещё раз",
		}
		err := b.sendMessage(msg, defaultChannel)
		if err != nil {
			return err
		}
	}

	userId := message.From.ID
	data := b.userData[userId]

	layout := "2006-01-02"
	parsed, err := time.Parse(layout, message.Text)
	if err != nil {
		return err
	}

	data.Deadline = parsed

	b.userData[userId] = data
	b.switcher.Next()
	return nil
}

func (b *Bot) handleWaitingId(message *tgbotapi.Message) {
	validate := validator.New()

	err := validate.Var(message.Text, "required,number")
	if err != nil {
		logrus.Errorf("failed to validate text: %v", err)
		return
	}

	data := b.userData[message.From.ID]
	data.Id, err = strconv.Atoi(message.Text)
	if err != nil {
		logrus.Errorf("failed to parse id: %v", err)
		return
	}

	msg := models.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите новое название вашего дз/записи или напишите /done",
	}

	err = b.sendMessage(msg, defaultChannel)
	if err != nil {
		logrus.Errorf("failed to send message: %v", err)
		return
	}

	b.userData[message.From.ID] = data
	b.switcher.Next()
}
