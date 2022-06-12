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

	// get genres
	query = `select 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.movie_id = $1`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.GenreID] = mg.Genre.GenreName
	}
	movie.MovieGenre = genres
	return &movie, nil
}

// Returns all movies and error if any
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
