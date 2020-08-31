package logchain

import (
	"github.com/vTCP-Foundation/observerd/core/e"
	"github.com/vTCP-Foundation/observerd/core/logchain"
)

type Handler struct {
	lc *logchain.Log
}

func New() (handler *Handler) {
	handler = &Handler{}
	return
}

func (handler *Handler) Run() (flow <-chan error) {
	errorsFlow := make(chan error)

	go func() {
		var err error
		defer e.ReportLocalError(&err, errorsFlow)

		// load storage or init new

		// generate next block.
	}()

	return errorsFlow
}

func (handler *Handler) initLogChain() (err error) {
	handler.lc = logchain.New()

}
