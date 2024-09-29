package converter

import (
	"fmt"
	"homework_bot/internal/domain"
	"strconv"
)

type ScheduleConv struct {
}

func NewScheduleConv() *ScheduleConv {
	return &ScheduleConv{}
}

func (c *ScheduleConv) subjectToText(subject domain.Subject) string {
	text := "_________________________________\n"
	timeSlots := map[int]string{
		8:  "1. " + subject.Title + "\n",
		10: "2. " + subject.Title + "\n",
		11: "3. " + subject.Title + "\n",
		13: "4. " + subject.Title + "\n",
		15: "5. " + subject.Title + "\n",
		16: "6. " + subject.Title + "\n",
		18: "7. " + subject.Title + "\n",
		20: "8. " + subject.Title + "\n",
	}

	text += timeSlots[subject.Start.Hour()]
	text += fmt.Sprintf("    %s:%s - %s:%s  %s\n\n", strconv.Itoa(subject.Start.Hour()),
		strconv.Itoa(subject.Start.Minute()), strconv.Itoa(subject.End.Hour()), strconv.Itoa(subject.End.Minute()), subject.Classroom)

	if subject.Teacher != "" {
		text += fmt.Sprintf("    %s", subject.Teacher)
		if subject.TeacherDegree != "" {
			text += fmt.Sprintf(" (%s)", subject.TeacherDegree)
		}
		text += "\n"
	}

	text += fmt.Sprintf("    %s\n", subject.PPSLoad)

	return text
}

func (c *ScheduleConv) ScheduleToText(schedule domain.Schedule) map[string]string {
	messages := make(map[string]string)

	for _, subject := range schedule.Subjects {
		info, ok := messages[subject.Start.Weekday().String()]
		dayOfWeek := map[string]string{
			"Monday":    "Понедельник",
			"Tuesday":   "Вторник",
			"Wednesday": "Среда",
			"Thursday":  "Четверг",
			"Friday":    "Пятница",
			"Saturday":  "Суббота",
		}

		message := c.subjectToText(subject)
		if !ok {
			info = fmt.Sprintf("%s %s (%s)\n", subject.Start.Format("02-01-2006"),
				dayOfWeek[subject.Start.Weekday().String()], subject.Subgroup)
		}
		info += message
		messages[subject.Start.Weekday().String()] = info
	}
	return messages
}