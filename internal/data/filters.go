package data

import "github.com/harshk200/greenlight/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string // for e.g. id, year, -year NOTE: -year means decending order ascending by default
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than 0")
	v.Check(f.Page < 10_000_000, "page_size", "must be smaller than max 10000000")
	v.Check(f.PageSize > 0, "page_size", "must be greater than 0")
	v.Check(f.PageSize > 0, "page_size", "maximum allowed value 100")

	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}
