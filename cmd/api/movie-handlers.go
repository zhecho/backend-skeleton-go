package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("Invalid id parameter"))
		app.errorJSON(w, err)
		return
	}
	app.logger.Println("Getting movie with id:", id)

	movie, err := app.models.DB.Get(id)

	// test data
	// movie := models.Movie{
	// ID:          id,
	// Title:       "The Shawshank Redemption",
	// Description: "Two imprisoned",
	// Year:        1994,
	// ReleaseDate: "1994-10-14",
	// Rating:      "R",
	// MPAARating:  "pg-13",
	// CreatedAt:   time.Now(),
	// UpdatedAt:   time.Now(),
	// Runtime:     "2h 22min",
	// Director:    "Frank Darabont",
	// Actors:      "Tim Robbins, Morgan Freeman, Bob Gunton",
	// }
	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
