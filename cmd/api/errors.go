package main

import (
	"fmt"
	"net/http"
)

// TODO: later include additional info about the request
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// NOTE: for message we are expecting any struct that we'll JSON encode
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	// NOTE: envelope struct is just an outer warpper type so we can specify
    // what we are sending like data, movie, etc.. (error in our case)
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// sends a 404 not found response in json format using the same convension as in the whole application
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
    message := "unable to process the record due to an edit conflict, please try again"
    app.errorResponse(w, r, http.StatusConflict, message)
}
