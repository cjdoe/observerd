package main

import (
	"github.com/rs/zerolog/log"
	"github.com/vTCP-Foundation/observerd/common/settings"
	"github.com/vTCP-Foundation/observerd/core"
)


func main() {
	var err error

	exitOnError := func() {
		if err != nil {
			log.Fatal().Err(err).Msg("Exit")
		}
	}

	err = settings.Load()
	exitOnError()

	err = <- core.New().Run()
	exitOnError()
}
