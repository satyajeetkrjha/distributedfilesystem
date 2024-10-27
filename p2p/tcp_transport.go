package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	// the mutex will protect the peers
	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NEWTCPTransport(ListenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: ListenAddress,
	}
}

func (t *TCPTransport) LisenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

}
