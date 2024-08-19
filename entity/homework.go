package entity

import "time"

type Homework struct {
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
	Tags        []string  `db:"tags"`
	CreatedAt   time.Time `db:"created_at"`
	Deadline    time.Time `db:"deadline"`
	UpdatedAt   time.Time `db:"updated_at"`
}
