package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/application/services"
	"homework_bot/internal/domain"
	"homework_bot/pkg/switcher"
)

type IBot interface {
	SendHomework(homework domain.HomeworkToGet, chatId int64, channel int) error
	SendSchedule(schedule domain.Schedule, chatId int64, channel int) error
	SendMessage(message domain.MessageToSend, channel int) error
	SendInputError(message *tgbotapi.Message) error
	GetUserStates() map[int64]string
	GetUserData() map[int64]domain.Homework
	SetUserStates(userStates map[int64]string)
	SetUserData(userData map[int64]domain.Homework)
	GetServices() *services.Service
	GetSwitcher() *switcher.Switcher
	GetBot() *tgbotapi.BotAPI
}

const (
	CommandStart            = "start"
	CommandAdd              = "add"
	CommandUpdate           = "update"
	CommandDelete           = "delete"
	CommandHelp             = "help"
	CommandGetAll           = "get_all"
	CommandGetOnWeek        = "get_on_week"
	CommandGetOnToday       = "get_on_today"
	CommandGetOnTomorrow    = "get_on_tomorrow"
	CommandGetOnDate        = "get_on_date"
	CommandGetOnId          = "get_on_id"
	CommandScheduleWeek     = "schedule_week"
	CommandScheduleDate     = "schedule_date"
	CommandScheduleToday    = "schedule_today"
	CommandScheduleTomorrow = "schedule_tomorrow"
)

const (
	WaitingId          = "WaitingId"
	WaitingName        = "WaitingName"
	WaitingDescription = "WaitingDescription"
	WaitingImages      = "WaitingImages"
	WaitingTags        = "WaitingTags"
	WaitingDeadline    = "WaitingDeadline"
)

const (
	DefaultChannel     = 0
	ChannelInformation = 2
	ChannelBot         = 5
)
