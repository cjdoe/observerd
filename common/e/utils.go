package e

import (
	"github.com/vTCP-Foundation/observerd/core/logger"
	"os"
)

func InterruptOnError(err error) {
	if err != nil {
		if logger.IsInitialized() {
			logger.Log.Error(err)
		}

		os.Exit(-1)
	}
}

func InterruptIfNil(arg interface{}) {
	if arg == nil {
		panic("nil argument is not allowed")
	}
}

func InterruptIfEmpty(arg string) {
	if arg == "" {
		panic("empty argument is not allowed")
	}
}

func ReportLocalError(err *error, flow chan<- error) {
	InterruptIfNil(err)

	if *err != nil {
		flow <- *err
	}
}
