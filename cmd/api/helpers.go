package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	idStr := params.ByName("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	// NOTE: the id will be a unique positive integer
	if err != nil || id < 1 {
		return 0, errors.New("Invalid id parameter")
	}

	return id, nil
}
