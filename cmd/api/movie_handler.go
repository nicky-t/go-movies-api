package main

import (
	"backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println("Invalid id paramater")
	}

	app.logger.Println("id is", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Some Title",
		Description: "Some Description",
		Year:        2022,
		ReleaseDate: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
		RunTime:     100,
		Raging:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
