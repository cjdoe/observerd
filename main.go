package main

import (
	"github.com/vTCP-Foundation/observerd/core"
	"github.com/vTCP-Foundation/observerd/core/e"
	"github.com/vTCP-Foundation/observerd/core/logger"
	"github.com/vTCP-Foundation/observerd/core/settings"
)

func main() {
	err := settings.Load()
	e.InterruptOnError(err)

	err = logger.Init()
	e.InterruptOnError(err)

	c, err := core.New()
	e.InterruptOnError(err)

	err = <-c.Run()
	logger.Log.Fatal(err)
}
