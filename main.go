package main

import (
	"fmt"
	"log"

	"github.com/ashish1401/dist_file_storage/p2p"
)

func main() {
	// The line `github.com/ashish1401/Distributed-File-Storage/p2p` is an import statement in Go that is
	// used to import a package named `p2p` from the `Distributed-File-Storage` repository under the user
	// `ashish1401` on GitHub. This allows the main program to access and use the functionalities provided
	// by the `p2p` package in the code.
	tr := p2p.NewTCPTransport(":3000")
	// log.Fatal(tr.ListenAndAccept())
	// The line `log.Fatal(tr.ListenAndAccept())` in the Go code snippet is calling the `Fatal` function
	// from the `log` package.

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World")
	select {}
}
