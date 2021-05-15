package core

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/vTCP-Foundation/observerd/core/logchain/producer"
	"sync"
)

type Core struct {
	producer *producer.Producer
}

func New() (core *Core) {
	core = &Core{
		producer: producer.New(),
	}

	return
}

func (c *Core) Run() (err error) {
	components := sync.WaitGroup{}
	criticalErrorsFlow := make(chan error)

	gracefullyStopAllComponents := func () {
		criticalErrorsFlow <- c.producer.Stop()
		components.Done()
	}

	components.Add(1)
	go func() {
		criticalErrorsFlow <- c.producer.Run()

		// If this line executes - previous line reported error
		// and all components must be stooped now.
		gracefullyStopAllComponents()
	}()

	// ...
	// Another component goes here.
	// WARN: Do not forget to update gracefullyStopAllComponents() method.
	// ...

	components.Wait()

	for {
		if len(criticalErrorsFlow) > 0 {
			log.Err(<- criticalErrorsFlow).Msg("Error caught on core's level")

		} else {
			break
		}
	}

	err = errors.New("core critical error")
	return
}
