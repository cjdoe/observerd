package core

import (
	logchain2 "github.com/vTCP-Foundation/observerd/core/handlers/logchain"
	"github.com/vTCP-Foundation/observerd/core/logger"
)

type Core struct {
	lc *logchain2.Handler
}

func New() (core *Core, err error) {
	core = &Core{
		lc: logchain2.New(),
	}

	return
}

func (c *Core) Run() (flow <-chan error) {
	errorsFlow := make(chan error)

	go func() {
		var fatalError error

		select {
		case fatalError = <-c.lc.Run():
		}

		logger.Log.Error(fatalError)

		// todo: stop gracefully
		//stopGracefully()

		errorsFlow <- fatalError
	}()

	return errorsFlow
}

//
//func runComponentsAndWatchForError() {
//	var fatalError error
//
//	select {
//	case fatalError = <-publicInterface.Run():
//		{
//		}
//	case fatalError = <-nodesInterface.Run():
//		{
//		}
//	case fatalError = <-transactionsHandler.Run():
//		{
//		}
//	}
//
//	logger.Log.Error(fatalError)
//	stopGracefully()
//}
//
//func stopGracefully() {
//	// WARN: The order is important here!
//	// Interfaces must be closed in the first order.
//	// This will prevent accepting of the new transactions,
//	// so there would not be transactions that would not be processed by the transactions handler.
//	publicInterface.Stop()
//	nodesInterface.Stop()
//
//	transactionsHandler.Stop()
//}
//
