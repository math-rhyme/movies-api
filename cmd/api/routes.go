package main 

import (
	"net/http"
    "github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// initializing a new httprouter router instance
	r := httprouter.New()

	// register the relevant methods, URL patterns and handler fucntios 
	r.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	r.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return r
}