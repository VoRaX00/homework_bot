package switchStatus

type Switcher struct {
	ISwitchBot
}

func NewSwitcher(statuses []string) *Switcher {
	return &Switcher{
		ISwitchBot: NewSwitcherBot(statuses),
	}
}
