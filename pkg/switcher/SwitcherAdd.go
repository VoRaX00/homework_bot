package switcher

type SwitcherAdd struct {
	currentStatus int
	statuses      []string
}

func NewSwitcherAdd(statuses []string) *SwitcherAdd {
	return &SwitcherAdd{
		currentStatus: -1,
		statuses:      statuses,
	}
}

func (s *SwitcherAdd) Next() {
	if s.currentStatus < len(s.statuses)-1 {
		s.currentStatus++
		return
	}
	s.currentStatus = -1
}

func (s *SwitcherAdd) Current() string {
	if s.currentStatus == -1 {
		return ""
	}
	return s.statuses[s.currentStatus]
}

func (s *SwitcherAdd) Previous() {
	switch s.currentStatus {
	case 0:
		s.currentStatus = -1
		break
	default:
		s.currentStatus--
	}
}

func (s *SwitcherAdd) IsActive() bool {
	return s.currentStatus >= 0
}
