package shared

import (
	"net"
)

type ServerStub struct {
	Identity
}

func (this *ServerStub) process(in, out net.Conn, methodId int) {

}
