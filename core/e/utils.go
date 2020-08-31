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

func InterruptOnNilArgument(arg interface{}) {
	if arg == nil {
		panic(arg)
	}
}

func ReportLocalError(err *error, flow chan<- error) {
	InterruptOnNilArgument(err)

	if *err != nil {
		flow <- *err
	}
}
