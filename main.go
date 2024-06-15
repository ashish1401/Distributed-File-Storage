package main

import (
	"fmt"

	"github.com/ashish1401/dist_fole_storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	fmt.Println("Hello World")
}
