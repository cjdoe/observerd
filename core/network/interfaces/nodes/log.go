package nodes

import (
	"github.com/vTCP-Foundation/observerd/core/network/interfaces"
	"net"
)

func (i *PublicInterface) logIngress(bytesReceived int, conn net.Conn) {
	r.log().Debug("[TX<=] ", bytesReceived, "B, ", conn.RemoteAddr())
}

func (i *PublicInterface) logEgress(bytesSent int, conn net.Conn) {
	r.log().Debug("[TX=>] ", bytesSent, "B, ", conn.RemoteAddr())
}

func (i *PublicInterface) log() *log.Entry {
	return log.WithFields(log.Fields{"prefix": "Network/GEO"})
}