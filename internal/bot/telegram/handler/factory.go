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
	userId := message.From.ID
	switch {
	case b.GetSwitcher().ISwitcherUpdate.Current(userId) == bot.WaitingId:
		return NewWaitingIdHandler()
	case b.GetSwitcher().ISwitcherUser.Current(userId) == bot.WaitingGroup:
		return NewAskGroupHandler()
	case b.GetSwitcher().ISwitcherAdd.Current(userId) == bot.WaitingName || b.GetSwitcher().ISwitcherUpdate.Current(userId) == bot.WaitingName:
		return NewWaitingNameHandler()
	case b.GetSwitcher().ISwitcherAdd.Current(userId) == bot.WaitingDescription || b.GetSwitcher().ISwitcherUpdate.Current(userId) == bot.WaitingDescription:
		return NewWaitingDescriptionHandler()
	case b.GetSwitcher().ISwitcherAdd.Current(userId) == bot.WaitingImages || b.GetSwitcher().ISwitcherUpdate.Current(userId) == bot.WaitingImages:
		return NewWaitingImageHandler()
	case b.GetSwitcher().ISwitcherAdd.Current(userId) == bot.WaitingTags ||
		b.GetSwitcher().ISwitcherUpdate.Current(userId) == bot.WaitingTags || b.GetSwitcher().ISwitcherGetTags.Current(userId) == bot.WaitingTags:
		return NewWaitingTagsHandler()
	case b.GetSwitcher().ISwitcherAdd.Current(userId) == bot.WaitingDeadline || b.GetSwitcher().ISwitcherUpdate.Current(userId) == bot.WaitingDeadline:
		return NewWaitingDeadlineHandler()
	default:
		if message.IsCommand() {
			return NewCommandHandler()
		}
		return NewMessageHandler()
	}
}
