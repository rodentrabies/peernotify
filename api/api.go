package api

import (
	"net/http"
	"net/url"

	"github.com/yurizhykin/peernotify/core"
)

const (
	API_PREFIX  = "" // "/api/"
	API_VERSION = "" // "v1"
)

func NewAPIServer(node *core.PeernotifyNode, addr string) http.Server {
	return http.Server{Addr: addr, Handler: newAPIHandler(node)}
}

type apiHandler struct {
	mux  map[string]func(http.ResponseWriter, *http.Request)
	node *core.PeernotifyNode
}

func (h *apiHandler) apiHandle(path string, handler func(http.ResponseWriter, *http.Request)) {
	h.HandleFunc(API_PREFIX+API_VERSION+path, handler)
}

func newAPIHandler(node *core.PeernotifyNode) *apiHandler {
	mux = map[string]func(http.ResponseWriter, *http.Request){
		"/register": h.handleRegister,
		"/verify":   h.handleVerify,
		"/forward":  h.handleForward,
	}
	h := &apiHandler{mux, node}
	return h
}

func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		return
	}
	// TODO: unify path string
	path := u.String()
	h.mux[path](w, r)
}

func apiErrorResponse(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (h *apiHandler) handleRegister(http.ResponseWriter, *http.Request) {
	apiErrorResponse(w, r)
}

func (h *apiHandler) handleVerify(http.ResponseWriter, *http.Request) {
	apiErrorResponse(w, r)
}

func (h *apiHandler) handleForward(http.ResponseWriter, *http.Request) {
	apiErrorResponse(w, r)
}
