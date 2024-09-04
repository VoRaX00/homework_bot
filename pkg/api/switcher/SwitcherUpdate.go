package switcher

type SwitcherUpdate struct {
	currentStatus int
	statuses      []string
}

func NewSwitcherUpdate(statuses []string) *SwitcherUpdate {
	return &SwitcherUpdate{
		currentStatus: -1,
		statuses:      statuses,
	}
}

func (s *SwitcherUpdate) Next() string {
	if s.currentStatus < len(s.statuses)-1 {
		return s.statuses[s.currentStatus]
	}
	s.currentStatus = -1
	return ""
}

func (s *SwitcherUpdate) Current() string {
	if s.currentStatus == -1 {
		return ""
	}
	return s.statuses[s.currentStatus]
}

func (s *SwitcherUpdate) Previous() string {
	switch s.currentStatus {
	case 0:
		s.currentStatus = -1
		return ""
	case 1:
		return ""
	default:
		s.currentStatus--
		return s.statuses[s.currentStatus]
	}
}

func (s *SwitcherUpdate) IsActive() bool {
	return s.currentStatus >= 0
}
