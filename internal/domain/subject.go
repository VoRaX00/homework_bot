package domain

import "time"

type Subject struct {
	Name        string
	BeginLesson time.Time
	EndLesson   time.Time
	TypeLesson  string
	Office      string
	Teacher     string
	Number      string
	Day         time.Time
}
