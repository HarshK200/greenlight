package data

import (
	"time"

	"github.com/harshk200/greenlight/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // this isn't relative for end-user hence use - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"` // movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "mustbe provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be longer than 500 bytes")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than or equal to 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genres")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 generes")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate generes")
}
