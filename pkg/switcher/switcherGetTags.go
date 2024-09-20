package switcher

type SwitcherGetTags struct {
	currentStatus int
	statuses      []string
}

func NewSwitcherGetTags(statuses []string) *SwitcherGetTags {
	return &SwitcherGetTags{
		currentStatus: -1,
		statuses:      statuses,
	}
}

func (s *SwitcherGetTags) Next() {
	if s.currentStatus < len(s.statuses)-1 {
		s.currentStatus++
		return
	}
	s.currentStatus = -1
}

func (s *SwitcherGetTags) Current() string {
	if s.currentStatus == -1 {
		return ""
	}
	return s.statuses[s.currentStatus]
}

func (s *SwitcherGetTags) Previous() {
	switch s.currentStatus {
	case 0:
		s.currentStatus = -1
		break
	default:
		s.currentStatus--
	}
}

func (s *SwitcherGetTags) IsActive() bool {
	return s.currentStatus >= 0
}
