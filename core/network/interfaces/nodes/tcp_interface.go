package nodes

import (
	"github.com/vTCP-Foundation/observerd/core/logger"
	"github.com/vTCP-Foundation/observerd/core/settings"
	"net"
)

type TCPInterface struct {
}

func NewNodesInterface() (i *TCPInterface, err error) {
	i = &TCPInterface{}
	return
}

func (i *TCPInterface) Run() (flow <-chan error) {
	errorsFlow := make(chan error)

	go func() {
		listener, err := net.Listen("tcp", settings.Conf.Interfaces.Clients.Interface())
		if err != nil {
			errorsFlow <- err
			return
		}

		// todo: this code never executes
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				logger.Log.Error(err)
				continue
			}

			handler, err := NewIncomingConnectionHandler(conn, i)
			if err != nil {
				logger.Log.Error(err)
				continue
			}

			go handler.Run()
		}
	}()

	return errorsFlow
}

func (i *TCPInterface) Stop() {

}

func (i *TCPInterface) Run() (errors <-chan error) {
	errorsFlow := make(chan error)

	return errorsFlow
}
