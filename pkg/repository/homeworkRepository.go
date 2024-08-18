package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"main.go/Entity"
)

type HomeworkRepository struct {
	db *sqlx.DB
}

func NewHomeworkRepository(db *sqlx.DB) *HomeworkRepository {
	return &HomeworkRepository{
		db: db,
	}
}

func (r *HomeworkRepository) Create(homework Entity.Homework) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return -1, err
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (name, description, image, deadline)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, homeworkTable)
	var homeworkId int
	row := tx.QueryRow(query, homework.Name, homework.Description, homework.Image, homework.Deadline)

	if err = row.Scan(&homeworkId); err != nil {
		return -1, err
	}

	query = fmt.Sprintf(`INSERT INTO %s (name) VALUES ($1)
					ON CONFLICT (name) DO NOTHING RETURNING id`, tagsTable)

	tagsId := make([]int, 0, len(homework.Tags))
	for _, tag := range homework.Tags {
		var tagId int
		err = tx.QueryRow(query, tag).Scan(&tagId)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return -1, err
		}

		if tagId > 0 {
			tagsId = append(tagsId, tagId)
		} else {
			getQuery := fmt.Sprintf(`SELECT id FROM %s WHERE name = $1`, tagsTable)
			err = tx.Get(&tagId, getQuery, tag)
			if err != nil {
				return -1, err
			}
			tagsId = append(tagsId, tagId)
		}
	}

	for _, tagId := range tagsId {
		query = fmt.Sprintf(`INSERT INTO %s (homework_id, tag_id) VALUES ($1, $2)`, homeworkTagsTable)
		_, err = tx.Exec(query, homeworkId, tagId)
		if err != nil {
			_ = tx.Rollback()
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return homeworkId, nil
}
