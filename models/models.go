package models

import (
	"database/sql"
	"time"
)

// Models is a wrapper for the database
type Models struct {
	DB DBModel
}

// Returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Type for movies
type Movie struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate string         `json:"release_date"`
	Genre       string         `json:"genre"`
	MovieGenre  map[int]string `json:"genres"`
	Poster      string         `json:"poster"`
	Plot        string         `json:"plot"`
	Rating      string         `json:"rating"`
	MPAARating  string         `json:"mpaa_rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Runtime     string         `json:"runtime"`
	Director    string         `json:"director"`
	Actors      string         `json:"actors"`
	Writers     string         `json:"writers"`
	Awards      string         `json:"awards"`
	BoxOffice   string         `json:"boxOffice"`
	Production  string         `json:"production"`
	Website     string         `json:"website"`
	IMDBID      string         `json:"imdbID"`
	Type        string         `json:"type"`
	Response    string         `json:"response"`
	Error       string         `json:"error"`
}

type Genre struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	GenreName string    `json:"genre_name"`
}

type MovieGenre struct {
	ID        int       `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	MovieID   int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
}
