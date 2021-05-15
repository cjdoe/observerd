package nodes

import (
	"github.com/vTCP-Foundation/observerd/common/settings"
	"net"
)

type TCPInterface struct {}

func NewNodesInterface() (i *TCPInterface, err error) {
	i = &TCPInterface{}
	return
}

func (i *TCPInterface) Run() (flow <-chan error) {
	errorsFlow := make(chan error)
	flow = errorsFlow

	listener, err := net.Listen("tcp", settings.Conf.Interfaces.Public.Interface())
	if err != nil {
		errorsFlow <- err
		return
	}

	defer func() {
		_ = listener.Close()
	}()

	// todo: add logger here

	for {
		conn, err := listener.Accept()
		if err != nil {
			errorsFlow <- err
			return
		}

		go i.handleConnection(conn, errorsFlow)
	}
}

func (i *TCPInterface) handleConnection(conn net.Conn, globalErrorsFlow chan<- error) {
	message, err := r.receiveData(conn)
	if err != nil {
		processError(err)
		return
	}

	request, e := v0.ParseRequest(message)
	if e != nil {
		processError(e)
		return
	}

	go r.handleRequest(conn, request, globalErrorsFlow)
}