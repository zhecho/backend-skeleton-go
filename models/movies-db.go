package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Returns one movie and error if any
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var movie Movie
	query := `select id, title, description, year, release_date, rating, runtime, mpaa_rating,
			  created_at, updated_at from movies where id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// Returns all movies and error if any
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
