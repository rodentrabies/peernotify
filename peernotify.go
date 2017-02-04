package main

import (
	"fmt"

	"github.com/yurizhykin/peernotify/api"
)

func main() {
	addr := ":2701"
	apiServer := api.NewAPIServer(addr)
	fmt.Printf("Starting peernotify server at %s\n", addr)
	apiServer.ListenAndServe()
}
