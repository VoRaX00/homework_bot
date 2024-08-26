package switcher

type Switcher struct {
	ISwitcherAdd
	ISwitcherUpdate
}

func NewSwitcher(statusesAdd []string, statusesUpdate []string) *Switcher {
	return &Switcher{
		ISwitcherAdd:    NewSwitcherAdd(statusesAdd),
		ISwitcherUpdate: NewSwitcherUpdate(statusesUpdate),
	}
}
