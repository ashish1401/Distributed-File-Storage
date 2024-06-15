package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPeer Represents the remote node over a TCP established connection
type TCPPeer struct {
	// conn is underlying connection of the peer
	conn net.Conn
	// (sender) if we dial and retrieve a connection -> outbound would be true
	// (receiver) if we accept and retrieve a connection -> inbound thus outbound would be false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	tp := TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
	return &tp
}

// struct -> pass by value, hence use of pointers, referencing and dereferencing
type TCPTransport struct {
	listenAddress string
	listener      net.Listener // accept, close, senders address

	mu    sync.RWMutex      //common practice to place mutexes above peers --> entity u want to protect
	peers map[net.Addr]Peer //net addr is an interface with Network() and Addr()

}

func NewTCPTransport(listenAddr string) *TCPTransport {
	t := &TCPTransport{
		listenAddress: listenAddr,
	}
	return t
}

// function for listening to the sender and curating a connection between sender and receiver
// each peer shall have the ability to ListenAndAccept
func (t *TCPTransport) ListenAndAccept() error {

	//initialization of error
	var err error
	// conn := net.Addr.Network()
	// save senders address ( .listener gives you - accept,close,sender address)
	t.listener, err = net.Listen("tcp", t.listenAddress) // ???
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept Error: %s", err)
		}

		go t.handleConn(conn)
	}

}

// connection handling to be called in a go routine
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("Incoming Connection : %v\n", peer)
}

// func Tset() {
// 	t, next := NewTCPTransport("199.10.22.44").listener.Accept()
// }
