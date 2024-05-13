package homework

import (
	"database/sql"
	"fmt"
	"time"
)

// Класс предстовляющий из себя домашнее задание с id, названием предмета и дедлайном
type Homework struct {
	subject  string
	content  string
	deadline time.Time
}

// конструктор для класса Homeworks
func NewHomework(subject string, content string, deadline time.Time) Homework {
	return Homework{subject: subject, content: content, deadline: deadline}
}

func (hw Homework) Subject() string {
	return hw.subject
}

func (hw Homework) Content() string {
	return hw.content
}

func (hw Homework) Deadline() time.Time {
	return hw.deadline
}

func (hw *Homework) SetSubject(subject string) {
	hw.subject = subject
}

func (hw *Homework) SetContent(content string) {
	hw.content = content
}

func (hw *Homework) SetDeadline(deadline time.Time) {
	hw.deadline = deadline
}

func (hw *Homework) AddHomeworkInDB(db *sql.DB) (bool, error) {
	req := fmt.Sprintf("INSERT INTO homework (subject, content, deadline) values('%v', '%v', '%v'))", hw.subject, hw.content, hw.deadline)
	result, err := db.Exec(req)
	if err != nil {
		return false, err
	}
	fmt.Println(result.RowsAffected())
	return true, nil
}
