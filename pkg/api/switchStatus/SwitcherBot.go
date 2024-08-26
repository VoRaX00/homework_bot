package switchStatus

type SwitcherBot struct {
	currentStatus int
	statuses      []string
}

func NewSwitcherBot(statuses []string) *SwitcherBot {
	return &SwitcherBot{
		currentStatus: -1,
		statuses:      statuses,
	}
}

func (s *SwitcherBot) Next() string {
	if s.currentStatus < len(s.statuses)-1 {
		return s.statuses[s.currentStatus]
	}

	s.currentStatus = -1
	return ""
}

func (s *SwitcherBot) Current() string {
	if s.currentStatus == -1 {
		return ""
	}
	return s.statuses[s.currentStatus]
}

func (s *SwitcherBot) Previous() string {
	if s.currentStatus == 0 {
		s.currentStatus = -1
		return ""
	}
	if s.currentStatus == -1 {
		return ""
	}

	s.currentStatus--
	return s.statuses[s.currentStatus]
}
