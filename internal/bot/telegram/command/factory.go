package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

const (
	commandStart         = "start"
	commandAdd           = "add"
	commandUpdate        = "update"
	commandDelete        = "delete"
	commandHelp          = "help"
	commandGetAll        = "get_all"
	commandGetOnWeek     = "get_on_week"
	commandGetOnToday    = "get_on_today"
	commandGetOnTomorrow = "get_on_tomorrow"
	commandGetOnDate     = "get_on_date"
	commandGetOnId       = "get_on_id"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) isAdmin(chatId int64) bool {
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

func (f *Factory) GetCommand(message *tgbotapi.Message) ICommand {
	if f.isAdmin(message.Chat.ID) {
		switch message.Command() {
		case commandAdd:
			return NewAddCommand()
		case commandUpdate:
			return NewUpdateCommand()
		case commandDelete:
			return NewDeleteCommand()
		}
	}

	switch message.Command() {
	case commandStart:
		return NewStartCommand()
	case commandHelp:
		return NewHelpCommand()
	case commandGetAll:
		return NewGetAllCommand()
	case commandGetOnId:
		return NewGetOnIdCommand()
	case commandGetOnWeek:
		return NewGetOnWeekCommand()
	case commandGetOnToday:
		return NewGetOnTodayCommand()
	case commandGetOnTomorrow:
		return NewGetOnTomorrowCommand()
	case commandGetOnDate:
		return NewGetOnDateCommand()
	default:
		return NewDefaultCommand()
	}
}
