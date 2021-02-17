package ec

import "errors"

var (
	ErrValidation = errors.New("validation error")
)

//
// Database related errors
//
var (
	ErrNoData = errors.New("no data fetched")
	ErrDBRead = errors.New("cant read from database")
)
