package api

import (
	"fmt"
	"net/http"
)

const (
	API_PREFIX  = "/api/"
	API_VERSION = "v1"
)

func apiHandle(path string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(API_PREFIX+API_VERSION+path, handler)
}

func APIServe(port int) {
	apiHandle("/register", handleRegister)
	apiHandle("/forward", handleForward)
	http.ListenAndServe(":"+string(port), nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register at %s", r.URL.Path)
}

func handleForward(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Forward at %s", r.URL.Path)
}
