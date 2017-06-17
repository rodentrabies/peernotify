package main

import (
	"log"

	"github.com/mrwhythat/peernotify/api"
	"github.com/mrwhythat/peernotify/core"
)

func main() {
	// Config
	addr := ":2701"
	store := "/home/whythat/.peernotify"

	// Setup
	node, err := core.NewPeernotifyNode(store)
	if err != nil {
		log.Fatalf("Error initializing peernotify node. Exiting...")
	}
	apiServer := api.NewAPIServer(node, addr)

	// Run
	log.Printf("Starting peernotify server at %s\n", addr)
	apiServer.ListenAndServe()
}
