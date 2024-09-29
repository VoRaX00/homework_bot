package switcher

type SwitcherUpdate struct {
	statuses      []string
	currentStatus map[int64]int
	users         map[int64]string
}

func NewSwitcherUpdate(statuses []string) *SwitcherUpdate {
	return &SwitcherUpdate{
		currentStatus: make(map[int64]int),
		users:         make(map[int64]string),
		statuses:      statuses,
	}
}

func (s *SwitcherUpdate) Next(id int64) {
	if s.currentStatus[id] < len(s.statuses)-1 {
		s.currentStatus[id]++
		return
	}
	s.currentStatus[id] = -1
}

func (s *SwitcherUpdate) Current(id int64) string {
	if s.currentStatus[id] == -1 {
		return ""
	}
	return s.statuses[s.currentStatus[id]]
}

func (s *SwitcherUpdate) Previous(id int64) {
	switch s.currentStatus[id] {
	case 0:
		s.currentStatus[id] = -1
		//return ""
	case 1:
		//return ""
	default:
		s.currentStatus[id]--
		//return s.statuses[s.currentStatus]
	}
}

func (s *SwitcherUpdate) IsActive(id int64) bool {
	return s.currentStatus[id] >= 0
}
