package main

import (
	"github.com/vTCP-Foundation/observerd/common/e"
	"github.com/vTCP-Foundation/observerd/common/settings"
	"github.com/vTCP-Foundation/observerd/core"
	"github.com/vTCP-Foundation/observerd/core/logger"
)

func main() {
	err := settings.Load()
	e.InterruptOnError(err)

	err = logger.Init()
	e.InterruptOnError(err)

	c := core.New()
	err = <-c.Run()
	logger.Log.Fatal(err)
}
