package switchStatus

type ISwitchBot interface {
	Current() string
	Next() string
	Previous() string
}
