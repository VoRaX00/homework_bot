package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) GetHandler(b bot.IBot, message *tgbotapi.Message) IHandler {
	switch {
	case b.GetSwitcher().ISwitcherUpdate.Current() == bot.WaitingId:
		return NewWaitingIdHandler()
	case b.GetSwitcher().ISwitcherAdd.Current() == bot.WaitingName || b.GetSwitcher().ISwitcherUpdate.Current() == bot.WaitingName:
		return NewWaitingNameHandler()
	case b.GetSwitcher().ISwitcherAdd.Current() == bot.WaitingDescription || b.GetSwitcher().ISwitcherUpdate.Current() == bot.WaitingDescription:
		return NewWaitingDescriptionHandler()
	case b.GetSwitcher().ISwitcherAdd.Current() == bot.WaitingImages || b.GetSwitcher().ISwitcherUpdate.Current() == bot.WaitingImages:
		return NewWaitingImageHandler()
	case b.GetSwitcher().ISwitcherAdd.Current() == bot.WaitingTags ||
		b.GetSwitcher().ISwitcherUpdate.Current() == bot.WaitingTags || b.GetSwitcher().ISwitcherGetTags.Current() == bot.WaitingTags:
		return NewWaitingTagsHandler()
	case b.GetSwitcher().ISwitcherAdd.Current() == bot.WaitingDeadline || b.GetSwitcher().ISwitcherUpdate.Current() == bot.WaitingDeadline:
		return NewWaitingDeadlineHandler()
	default:
		if message.IsCommand() {
			return NewCommandHandler()
		}
		return NewMessageHandler()
	}
}
