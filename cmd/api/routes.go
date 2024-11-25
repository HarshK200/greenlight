package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// routes
	router.GET("/v1/healthcheck", app.healthcheckHandler)
	//movies routes
    router.GET("/v1/movies", app.listMovieHandler)
	router.POST("/v1/movies", app.createMovieHandler)
	router.GET("/v1/movies/:id", app.showMovieHandler)
	router.PATCH("/v1/movies/:id", app.updateMovieHandler)
	router.DELETE("/v1/movies/:id", app.deleteMovieHandler)

	return router
}
