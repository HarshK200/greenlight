package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
    ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	DB     *sql.DB
	Movies MovieModel
}

// constructor
func NewModels(db *sql.DB) Models {
	return Models{
		DB:     db,
		Movies: MovieModel{DB: db},
	}
}
