package main

import (
	"fmt"
	"net/http"

	"github.com/yurizhykin/peernotify/api"
)

func main() {
	fmt.Println("Starting peernotify server...")
	api.APIServe(2701)
	http.ListenAndServe(":2701", nil)
}
