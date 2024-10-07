package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Blog struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tag       []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type BlogModel struct {
	DB *sql.DB
}

func (b *BlogModel) Insert(blog *Blog) error {
	query := `
	INSERT INTO blog(title, content, tag, updated_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at`

	args := []interface{}{blog.Title, blog.Content, pq.Array(blog.Tag), blog.UpdatedAt}

	return b.DB.QueryRow(query, args...).Scan(&blog.ID, &blog.CreatedAt)
}

func (b *BlogModel) Update(blog *Blog) error {
	query := `
	UPDATE blog
	SET title=$1, content=$2, tag=$3, updated_at=CURRENT_TIMESTAMP
	WHERE id=$4
	RETURNING updated_at`

	args := []interface{}{
		blog.Title,
		blog.Content,
		pq.Array(blog.Tag),
		blog.ID,
	}

	return b.DB.QueryRow(query, args...).Scan(&blog.UpdatedAt)

}

func (b *BlogModel) Get(id int64) (*Blog, error) {
	if id < 1 {
		return nil, ErrorRecordNotFound
	}

	query := `
	SELECT id, title, content, tag, created_at, updated_at
	FROM blog
	WHERE id=$1`

	var blog Blog

	err := b.DB.QueryRow(query, id).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		pq.Array(&blog.Tag),
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrorRecordNotFound
		default:
			return nil, err
		}
	}

	return &blog, nil
}

func (b *BlogModel) Delete(id int64) error {
	if id < 1 {
		return ErrorRecordNotFound
	}

	query := `
	DELETE
	FROM blog
	WHERE id=$1`

	result, err := b.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorRecordNotFound
	}

	return nil
}
