package ec

import (
	"errors"
	"github.com/vTCP-Foundation/observerd/core/logger"
	"os"
)

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

func InterruptOnError(err error) {
	if err != nil {
		if logger.IsInitialized() {
			logger.Log.Error(err)
		}

		os.Exit(-1)
	}
}
