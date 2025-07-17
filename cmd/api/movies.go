package main

import (
	"fmt"
	"net/http"
)

// handler for `POST /v1/movies` endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// get the value of the "id" parameter from the slice ((stored in the request context)).
	// if the parameter can't be converted or < 1, the ID is invalid

	// params := httprouter.ParamsFromContext(r.Context())
	// id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
