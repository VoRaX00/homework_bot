package switcher

type Switcher struct {
	ISwitcherAdd
	ISwitcherUpdate
}

func (s *Switcher) Next() {
	if s.ISwitcherAdd.IsActive() {
		s.ISwitcherAdd.Next()
	} else {
		s.ISwitcherUpdate.Next()
	}
}

func NewSwitcher(statusesAdd []string, statusesUpdate []string) *Switcher {
	return &Switcher{
		ISwitcherAdd:    NewSwitcherAdd(statusesAdd),
		ISwitcherUpdate: NewSwitcherUpdate(statusesUpdate),
	}
}
