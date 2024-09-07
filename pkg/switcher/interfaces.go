package switcher

type ISwitcher interface {
	Current() string
	Next() string
	Previous() string
	IsActive() bool
}

type ISwitcherAdd interface {
	ISwitcher
}

type ISwitcherUpdate interface {
	ISwitcher
}
