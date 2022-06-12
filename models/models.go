package models

import "time"

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate string       `json:"release_date"`
	Genre       string       `json:"genre"`
	MovieGenre  []MovieGenre `json:"-"` // ignore that json field
	Poster      string       `json:"poster"`
	Plot        string       `json:"plot"`
	Rating      string       `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Runtime     string       `json:"runtime"`
	Director    string       `json:"director"`
	Actors      string       `json:"actors"`
	Writers     string       `json:"writers"`
	Awards      string       `json:"awards"`
	BoxOffice   string       `json:"boxOffice"`
	Production  string       `json:"production"`
	Website     string       `json:"website"`
	IMDBID      string       `json:"imdbID"`
	Type        string       `json:"type"`
	Response    string       `json:"response"`
	Error       string       `json:"error"`
}

type Genre struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GenreName string    `json:"genre_name"`
}

type MovieGenre struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	MovieID   int       `json:"movie_id"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
}
