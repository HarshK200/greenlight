package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // this isn't relative for end-user hence use - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"` // movie runtime (in minutes)
	Geners    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie() {

}
