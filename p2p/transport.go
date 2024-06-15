package p2p

// Peer is an interface that represents the remote node
// Entity in the computer network
type Peer interface{}

// Transport is anything that handles the communication
// Between the nodes in the computer network
// Channel of communication between two peers
// Can be of form TCP, UDP, WebSocket
type Transport interface {
	ListenAndAccept() error
}
