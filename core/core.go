package core

import "github.com/vTCP-Foundation/observerd/core/logchain/producer"

type Core struct {
	producer *producer.Producer
}

func New() (core *Core) {
	core = &Core{
		producer: producer.New(),
	}

	return
}

func (c *Core) Run() (errorsFlow <-chan error) {
	flow := make(chan error)

	go func() {
		var err error

		select {
		case err = <-c.producer.Run():
			{
			}

		}

		flow <- err
	}()

	return flow
}
