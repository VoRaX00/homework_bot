package entity

import "time"

type Homework struct {
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Images      []string  `db:"images"`
	Tags        []string  `db:"tags"`
	CreatedAt   time.Time `db:"create_at"`
	Deadline    time.Time `db:"deadline"`
	UpdatedAt   time.Time `db:"update_at"`
}