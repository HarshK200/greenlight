package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/harshk200/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	// NOTE: the id will be a unique positive integer
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "dezznuts",
		Runtime:   102,
		Geners:    []string{"sci-fi", "horror"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and couldn't provess your request", http.StatusInternalServerError)
	}
}
