package switcher

type SwitcherAdd struct {
	statuses      []string
	currentStatus map[int64]int
	users         map[int64]string
}

func NewSwitcherAdd(statuses []string) *SwitcherAdd {
	return &SwitcherAdd{
		currentStatus: make(map[int64]int),
		users:         make(map[int64]string),
		statuses:      statuses,
	}
}

func (s *SwitcherAdd) Next(id int64) {
	if s.currentStatus[id] < len(s.statuses)-1 {
		s.currentStatus[id]++
		return
	}
	s.currentStatus[id] = -1
}

func (s *SwitcherAdd) Current(id int64) string {
	val, ok := s.currentStatus[id]
	if !ok || val == -1 {
		return ""
	}

	return s.statuses[s.currentStatus[id]]
}

func (s *SwitcherAdd) Previous(id int64) {
	val, ok := s.currentStatus[id]
	if !ok {
		return
	}

	switch val {
	case 0:
		s.currentStatus[id] = -1
		break
	default:
		s.currentStatus[id]--
	}
}

func (s *SwitcherAdd) IsActive(id int64) bool {
	_, ok := s.currentStatus[id]
	if ok {
		return s.currentStatus[id] >= 0
	}
	return false
}
