package api

import (
	"fmt"
	"net/http"
)

const (
	API_PREFIX  = "/api/"
	API_VERSION = "v1"
)

func NewAPIServer(addr string) http.Server {
	return http.Server{Addr: addr, Handler: newAPIHandler()}
}

type apiHandler struct {
	*http.ServeMux
}

func (h *apiHandler) apiHandle(path string, handler func(http.ResponseWriter, *http.Request)) {
	h.HandleFunc(API_PREFIX+API_VERSION+path, handler)
}

func newAPIHandler() *apiHandler {
	h := &apiHandler{http.NewServeMux()}
	h.apiHandle("/register", handleRegister)
	h.apiHandle("/verify", handleVerify)
	h.apiHandle("/forward", handleForward)
	return h
}

// Handlers
func handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register at %s", r.URL.Path)
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Verify at %s", r.URL.Path)
}

func handleForward(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Forward at %s", r.URL.Path)
}
