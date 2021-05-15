package nodes

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/vTCP-Foundation/observerd/core/common"
	"github.com/vTCP-Foundation/observerd/core/common/errors"
	"net"
)

func (i *TCPInterface) receiveData(conn net.Conn) (data []byte, err error) {
	reader := bufio.NewReader(conn)

	messageSizeBinary := []byte{0, 0, 0, 0}
	bytesRead, err := reader.Read(messageSizeBinary)
	if err != nil {
		err = fmt.Errorf("can't read message size header from tcp socket: %w", err)
		return
	}

	if bytesRead != len(messageSizeBinary) {
		err = fmt.Errorf("can't read message size header from tcp socket: %w", err)
		return
	}

	messageSize := binary.BigEndian.Uint64(messageSizeBinary)
	if messageSize > MaxMessageSize || messageSize < MinMessageSize {
		e = errors.AppendStackTrace(errors.InvalidDataFormat)
		return
	}

	var offset uint32 = 0
	data = make([]byte, messageSize, messageSize)
	for {
		bytesReceived, err := reader.Read(data[offset:])
		if err != nil {
			return nil, errors.AppendStackTrace(err)
		}
		if bytesReceived == 0 {
			return nil, errors.AppendStackTrace(errors.NoData)
		}

		offset += uint32(bytesReceived)
		if offset == messageSize {
			r.logIngress(len(data), conn)
			return data, nil
		}
	}
}
