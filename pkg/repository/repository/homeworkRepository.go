package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"main.go/pkg/entity"
	"main.go/pkg/repository/config"
	"time"
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

	// Вставка данных в таблицу homework
	query := fmt.Sprintf(`
		INSERT INTO %s (name, description, images, deadline)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, config.HomeworkTable)

	var homeworkId int
	row := tx.QueryRow(query, homework.Name, homework.Description, pq.Array(homework.Images), homework.Deadline)

	// Если ошибка, выполняем откат транзакции
	if err = row.Scan(&homeworkId); err != nil {
		_ = tx.Rollback()
		return -1, err
	}

	// Вставка тегов и связывание их с заданием
	query = fmt.Sprintf(`INSERT INTO %s (name) VALUES ($1)
					ON CONFLICT (name) DO NOTHING RETURNING id`, config.TagsTable)

	tagsId := make([]int, 0, len(homework.Tags))
	for _, tag := range homework.Tags {
		var tagId int
		err = tx.QueryRow(query, tag).Scan(&tagId)
		if err != nil && err != sql.ErrNoRows { // Исправлено условие на ErrNoRows
			_ = tx.Rollback()
			return -1, err
		}

		if tagId == 0 {
			getQuery := fmt.Sprintf(`SELECT id FROM %s WHERE name = $1`, config.TagsTable)
			err = tx.Get(&tagId, getQuery, tag)
			if err != nil {
				_ = tx.Rollback()
				return -1, err
			}
		}
		tagsId = append(tagsId, tagId)
	}

	// Связывание тегов с заданием
	for _, tagId := range tagsId {
		query = fmt.Sprintf(`INSERT INTO %s (homework_id, tag_id) VALUES ($1, $2)`, config.HomeworkTagsTable)
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

func (r *HomeworkRepository) GetByTags(tags []string) ([]entity.HomeworkToGet, error) {
	query := fmt.Sprintf(`SELECT h.* 
		FROM %s h
		JOIN %s ht ON h.id = ht.homework_id
		JOIN %s t ON ht.tag_id = t.id
		WHERE t.name = ANY($1)
		GROUP BY h.id
		HAVING COUNT(DISTINCT t.name) = $2;`, config.HomeworkTable, config.HomeworkTagsTable, config.TagsTable)

	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query, tags, len(tags))

	return homeworks, err
}

func (r *HomeworkRepository) GetByName(name string) ([]entity.HomeworkToGet, error) {
	query := fmt.Sprintf(`SELECT h.name, h.description, h.image, h.created_at, h.deadline, h.updated_at, ARRAY_AGG(t.name) AS %s
		FROM %s h 
		LEFT JOIN %s ht 
		ON h.id = ht.homework_id
		LEFT JOIN %s t 
		ON ht.tag_id = t.id WHERE h.name = $1 GROUP BY h.id;`, config.TagsTable, config.HomeworkTable, config.HomeworkTagsTable, config.TagsTable)
	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query, name)
	return homeworks, err
}

func (r *HomeworkRepository) GetById(id int) (entity.HomeworkToGet, error) {
	query := fmt.Sprintf(`
		SELECT h.name, h.description, h.image, h.created_at, h.deadline, h.updated_at, ARRAY_AGG(t.name) AS tags
		FROM %s h 
		LEFT JOIN %s ht ON h.id = ht.homework_id
		LEFT JOIN %s t ON ht.tag_id = t.id 
		WHERE h.id = $1 
		GROUP BY h.id;`, config.HomeworkTable, config.HomeworkTagsTable, config.TagsTable)

	var homework entity.HomeworkToGet
	err := r.db.Get(&homework, query, id)
	return homework, err
}

func (r *HomeworkRepository) GetByWeek() ([]entity.HomeworkToGet, error) {
	query := fmt.Sprintf(`
		SELECT h.name, h.description, h.image, h.created_at, h.deadline, h.updated_at
		FROM %s h
		WHERE h.deadline >= DATE_TRUNC('week', NOW())
		AND h.deadline < DATE_TRUNC('week', NOW()) + INTERVAL '1 week';`, config.HomeworkTable)

	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query)

	return homeworks, err
}

func (r *HomeworkRepository) GetAll() ([]entity.HomeworkToGet, error) {
	query := `SELECT h.id, h.name, h.description, h.images, h.create_at, h.deadline, h.update_at, 
        COALESCE(array_agg(t.name ORDER BY t.name), '{}') AS tags
    FROM 
        homework h
    LEFT JOIN 
        homeworks_tags ht ON h.id = ht.homework_id
    LEFT JOIN 
        tags t ON ht.tag_id = t.id
    GROUP BY 
        h.id, h.name, h.description, h.images, h.create_at, h.deadline, h.update_at;
    `

	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query)

	return homeworks, err
}

func (r *HomeworkRepository) GetByToday() ([]entity.HomeworkToGet, error) {
	query := `SELECT h.id, h.name, h.description, h.images, h.create_at, h.deadline, h.update_at, 
    COALESCE(array_agg(t.name ORDER BY t.name), '{}') AS tags
	FROM 
		homework h
	LEFT JOIN 
		homeworks_tags ht ON h.id = ht.homework_id
	LEFT JOIN 
		tags t ON ht.tag_id = t.id
	WHERE 
		h.deadline >= CURRENT_DATE 
		AND h.deadline < CURRENT_DATE + INTERVAL '1 day'
	GROUP BY 
		h.id, h.name, h.description, h.update_at, h.deadline, h.create_at, h.images;
		`

	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query)

	return homeworks, err
}

func (r *HomeworkRepository) GetByTomorrow() ([]entity.HomeworkToGet, error) {
	query := `SELECT h.id, h.name, h.description, h.images, h.create_at, h.deadline, h.update_at, 
    COALESCE(array_agg(t.name ORDER BY t.name), '{}') AS tags
	FROM 
		homework h
	LEFT JOIN 
		homeworks_tags ht ON h.id = ht.homework_id
	LEFT JOIN 
		tags t ON ht.tag_id = t.id
	WHERE 
		h.deadline >= CURRENT_DATE + INTERVAL '1 day'
		AND h.deadline < CURRENT_DATE + INTERVAL '2 days'
	GROUP BY 
		h.id, h.name, h.description, h.update_at, h.deadline, h.create_at, h.images;
		`

	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query)

	return homeworks, err
}

func (r *HomeworkRepository) GetByDate(date time.Time) ([]entity.HomeworkToGet, error) {
	formattedDate := date.Format("2006-01-02")

	query := `
    SELECT 
        h.id, 
        h.name, 
        h.description, 
        h.images, 
        h.create_at, 
        h.deadline, 
        h.update_at, 
        COALESCE(array_agg(t.name ORDER BY t.name), '{}') AS tags
    FROM 
        homework h
    LEFT JOIN 
        homeworks_tags ht ON h.id = ht.homework_id
    LEFT JOIN 
        tags t ON ht.tag_id = t.id
    WHERE 
        h.deadline >= $1::date 
        AND h.deadline < ($1::date + INTERVAL '1 day')
    GROUP BY 
        h.id, h.name, h.description, h.update_at, h.deadline, h.create_at, h.images;
    `

	var homeworks []entity.HomeworkToGet
	err := r.db.Select(&homeworks, query, formattedDate)

	return homeworks, err
}

func (r *HomeworkRepository) Update(homeworkToUpdate entity.HomeworkToUpdate) (entity.Homework, error) {
	query := "UPDATE " + config.HomeworkTable + " SET "
	var args []interface{}
	argIndex := 1

	if homeworkToUpdate.Name != nil {
		query += fmt.Sprintf("name = $%d, ", argIndex)
		args = append(args, *homeworkToUpdate.Name)
		argIndex++
	}

	if homeworkToUpdate.Description != nil {
		query += fmt.Sprintf("description = $%d, ", argIndex)
		args = append(args, *homeworkToUpdate.Description)
		argIndex++
	}

	if homeworkToUpdate.Images != nil {
		query += fmt.Sprintf("image = $%d::TEXT[], ", argIndex)
		args = append(args, pq.Array(*homeworkToUpdate.Images))
		argIndex++
	}

	if homeworkToUpdate.Deadline != nil {
		query += fmt.Sprintf("deadline = $%d, ", argIndex)
		args = append(args, *homeworkToUpdate.Deadline)
		argIndex++
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING *;", argIndex)
	args = append(args, homeworkToUpdate.Id)

	var updatedHomework entity.Homework
	err := r.db.Get(&updatedHomework, query, args...)
	if err != nil {
		return entity.Homework{}, err
	}

	if homeworkToUpdate.Tags != nil {
		deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE homework_id = $1;", config.HomeworkTagsTable)
		_, err = r.db.Exec(deleteQuery, homeworkToUpdate.Id)
		if err != nil {
			return entity.Homework{}, err
		}

		for _, tag := range *homeworkToUpdate.Tags {
			insertQuery := fmt.Sprintf(`
				INSERT INTO %s (homework_id, tag_id)
				VALUES ($1, (SELECT id FROM %s WHERE name = $2));`, config.HomeworkTagsTable, config.TagsTable)
			_, err = r.db.Exec(insertQuery, homeworkToUpdate.Id, tag)
			if err != nil {
				return entity.Homework{}, err
			}
		}
	}

	return updatedHomework, nil
}

func (r *HomeworkRepository) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM homeworks_tags WHERE homework_id = $1`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM homework WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
