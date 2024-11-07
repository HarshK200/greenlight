package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/harshk200/greenlight/internal/data"
	"github.com/harshk200/greenlight/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// NOTE: a use n' throw validator
	v := validator.New()

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	data.ValidateMovie(v, movie)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	// NOTE: the id will be a unique positive integer
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// HACK: stubs / mock data
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "dezznuts",
		Runtime:   102,
		Genres:    []string{"sci-fi", "horror"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
