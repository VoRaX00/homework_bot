package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"main.go/entity"
)

type HomeworkRepository struct {
	db *sqlx.DB
}

func NewHomeworkRepository(db *sqlx.DB) *HomeworkRepository {
	return &HomeworkRepository{
		db: db,
	}
}

func (r *HomeworkRepository) Create(homework entity.Homework) (int, error) {
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

func (r *HomeworkRepository) GetByTags(tags []string) ([]entity.Homework, error) {
	query := fmt.Sprintf(`SELECT h.* 
		FROM %s h
		JOIN %s ht ON h.id = ht.homework_id
		JOIN %s t ON ht.tag_id = t.id
		WHERE t.name = ANY($1)
		GROUP BY h.id
		HAVING COUNT(DISTINCT t.name) = $2;`, homeworkTable, homeworkTagsTable, tagsTable)

	var homeworks []entity.Homework
	err := r.db.Select(&homeworks, query, tags, len(tags))

	return homeworks, err
}

func (r *HomeworkRepository) GetByName(name string) ([]entity.Homework, error) {
	query := fmt.Sprintf(`SELECT h.name, h.description, h.image, h.created_at, h.deadline, h.updated_at, ARRAY_AGG(t.name) AS %s
		FROM %s h 
		LEFT JOIN %s ht 
		ON h.id = ht.homework_id
		LEFT JOIN %s t 
		ON ht.tag_id = t.id WHERE h.name = $1 GROUP BY h.id;`, tagsTable, homeworkTable, homeworkTagsTable, tagsTable)
	var homeworks []entity.Homework
	err := r.db.Select(&homeworks, query, name)
	return homeworks, err
}

func (r *HomeworkRepository) GetById(id int) (entity.Homework, error) {
	query := fmt.Sprintf(`SELECT h.name, h.description, h.image, h.created_at, h.deadline, h.updated_at, ARRAY_AGG(t.name) AS %s
		FROM %s h 
		LEFT JOIN %s ht 
		ON h.id = ht.homework_id
		LEFT JOIN %s t 
		ON ht.tag_id = t.id WHERE h.id = $1 GROUP BY h.id;`, tagsTable, homeworkTable, homeworkTagsTable, tagsTable)
	var homeworks entity.Homework
	err := r.db.Select(&homeworks, query, id)
	return homeworks, err
}

func (r *HomeworkRepository) GetByWeek() ([]entity.Homework, error) {
	query := fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE deadline >= DATE_TRUNC('week', NOW())
		AND deadline < DATE_TRUNC('week', NOW()) + INTERVAL '1 week';`, homeworkTable)

	var homeworks []entity.Homework
	err := r.db.Select(&homeworks, query)

	return homeworks, err
}

func (r *HomeworkRepository) GetAll() ([]entity.Homework, error) {
	query := fmt.Sprintf(`SELECT h.name, h.description, h.image, h.created_at, h.deadline, h.updated_at, ARRAY_AGG(t.name) AS %s
		FROM %s h 
		LEFT JOIN %s ht 
		ON h.id = ht.homework_id
		LEFT JOIN %s t 
		ON ht.tag_id = t.id GROUP BY h.id;`, tagsTable, homeworkTable, homeworkTagsTable, tagsTable)
	var homeworks []entity.Homework
	err := r.db.Select(&homeworks, query)

	return homeworks, err
}
