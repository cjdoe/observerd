package main

import (
	"github.com/vTCP-Foundation/observerd/common/settings"
	"github.com/vTCP-Foundation/observerd/core"
	"github.com/vTCP-Foundation/observerd/core/ec"
	"github.com/vTCP-Foundation/observerd/core/logger"
)

func main() {
	err := settings.Load()
	ec.InterruptOnError(err)

	err = logger.Init()
	ec.InterruptOnError(err)

	c := core.New()
	err = <-c.Run()
	logger.Log.Fatal(err)
}
