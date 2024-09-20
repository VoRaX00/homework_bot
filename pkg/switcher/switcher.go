package switcher

type Switcher struct {
	ISwitcherAdd
	ISwitcherUpdate
	ISwitcherGetTags
}

func (s *Switcher) Next() {
	if s.ISwitcherAdd.IsActive() {
		s.ISwitcherAdd.Next()
	} else if s.ISwitcherUpdate.IsActive() {
		s.ISwitcherUpdate.Next()
	} else {
		s.ISwitcherGetTags.Next()
	}
}

func NewSwitcher(statusesAdd []string, statusesUpdate []string, statusesGetTags []string) *Switcher {
	return &Switcher{
		ISwitcherAdd:     NewSwitcherAdd(statusesAdd),
		ISwitcherUpdate:  NewSwitcherUpdate(statusesUpdate),
		ISwitcherGetTags: NewSwitcherGetTags(statusesGetTags),
	}
}
