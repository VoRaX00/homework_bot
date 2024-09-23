package models

import "time"

type Subject struct {
	name        string    `db:"name"`
	place       string    `db:"place"`
	fromTime    time.Time `db:"from_time"`
	toTime      time.Time `db:"to_time"`
	teacher     string    `db:"teacher"`
	numberGroup string    `db:"number_group"`
}
