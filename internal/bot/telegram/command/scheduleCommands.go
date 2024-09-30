package command

import (
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"time"
)

type ScheduleWeekCommand struct{}

func NewScheduleWeekCommand() *ScheduleWeekCommand {
	return &ScheduleWeekCommand{}
}

func (c *ScheduleWeekCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	schedule := b.GetServices().GetOnWeek()
	err := b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleDayCommand struct{}

func NewScheduleDayCommand() *ScheduleDayCommand {
	return &ScheduleDayCommand{}
}

func (c *ScheduleDayCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	validate := validator.New()

	err := validate.RegisterValidation("customDate", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		_, err := time.Parse("2006-01-02", dateStr)
		return err == nil
	})
	if err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", message.Text)
	if err != nil {
		return err
	}

	schedule := b.GetServices().GetOnDate(date)
	err = b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleTodayCommand struct{}

func NewScheduleTodayCommand() *ScheduleTodayCommand {
	return &ScheduleTodayCommand{}
}

func (c *ScheduleTodayCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	schedule := b.GetServices().GetOnToday()
	err := b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleTomorrowCommand struct{}

func NewScheduleTomorrowCommand() *ScheduleTomorrowCommand {
	return &ScheduleTomorrowCommand{}
}

func (c *ScheduleTomorrowCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	schedule := b.GetServices().GetOnTomorrow()
	err := b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}
