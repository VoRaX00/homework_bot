package Entity

import "time"

type Homework struct {
	Name        string
	Description string
	Image       string
	Tags        []string
	CreatedAt   time.Time
	Deadline    time.Time
	UpdatedAt   time.Time
}
