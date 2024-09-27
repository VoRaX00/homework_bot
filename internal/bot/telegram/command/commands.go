package command

import (
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
	"strconv"
	"strings"
	"time"
)

type StartCommand struct{}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	textStart := "Привет! Меня зовут Биба, я буду твоим помошником для получения домашек и иных новосотей!"
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textStart,
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type AddCommand struct{}

func NewAddCommand() *AddCommand {
	return &AddCommand{}
}

func (c *AddCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	b.GetSwitcher().ISwitcherAdd.Next()
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите название домашней работы/записи",
	}

	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type UpdateCommand struct{}

func NewUpdateCommand() *UpdateCommand {
	return &UpdateCommand{}
}

func (c *UpdateCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	b.GetSwitcher().ISwitcherUpdate.Next()
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите Id вашей записи",
	}

	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type DeleteCommand struct{}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{}
}

func (c *DeleteCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	words := strings.Split(message.Text, " ")
	if len(words) != 2 {
		return b.SendInputError(message)
	}

	id, err := strconv.Atoi(words[1])
	if err != nil {
		return b.SendInputError(message)
	}

	err = b.GetServices().Delete(id)
	if err != nil {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Ошибка удаления",
		}
		_ = b.SendMessage(msg, bot.DefaultChannel)
		return err
	}

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Запись успешно удалена",
	}
	err = b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type GetAllCommand struct{}

func NewGetAllCommand() *GetAllCommand {
	return &GetAllCommand{}
}

func (c *GetAllCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetAll()

	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetOnWeekCommand struct{}

func NewGetOnWeekCommand() *GetOnWeekCommand {
	return &GetOnWeekCommand{}
}

func (c *GetOnWeekCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetByWeek()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}

	return nil
}

type GetOnIdCommand struct{}

func NewGetOnIdCommand() *GetOnIdCommand {
	return &GetOnIdCommand{}
}

func (c *GetOnIdCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	words := strings.Split(message.Text, " ")
	if len(words) != 2 {
		return b.SendInputError(message)
	}

	id, err := strconv.Atoi(words[1])
	if err != nil {
		return err
	}

	homework, err := b.GetServices().GetById(id)
	if err != nil {
		return err
	}

	err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
	return err
}

type GetOnTodayCommand struct{}

func NewGetOnTodayCommand() *GetOnTodayCommand {
	return &GetOnTodayCommand{}
}

func (c *GetOnTodayCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetByToday()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetOnTomorrowCommand struct{}

func NewGetOnTomorrowCommand() *GetOnTomorrowCommand {
	return &GetOnTomorrowCommand{}
}

func (c *GetOnTomorrowCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetByTomorrow()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetOnDateCommand struct{}

func NewGetOnDateCommand() *GetOnDateCommand {
	return &GetOnDateCommand{}
}

func (c *GetOnDateCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	words := strings.Split(message.Text, " ")
	if len(words) != 2 {
		return b.SendInputError(message)
	}

	date, err := time.Parse(time.DateOnly, words[1])
	if err != nil {
		return err
	}

	homeworks, err := b.GetServices().GetByDate(date)
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type HelpCommand struct{}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (c *HelpCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	textHelp := "Инструкция пользования Бибой:"
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textHelp,
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type DefaultCommand struct{}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (c *DefaultCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Я не знаком с такой командой :(",
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

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
