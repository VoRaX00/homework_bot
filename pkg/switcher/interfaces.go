package switcher

type ISwitcher interface {
	Current() string
	Next()
	Previous()
	IsActive() bool
}

type ISwitcherAdd interface {
	ISwitcher
}

type ISwitcherUpdate interface {
	ISwitcher
}

type ISwitcherGetTags interface {
	ISwitcher
}
