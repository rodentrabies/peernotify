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
	node *core.PeernotifyNode
}

func (h *apiHandler) apiHandle(path string, handler func(http.ResponseWriter, *http.Request)) {
	h.HandleFunc(API_PREFIX+API_VERSION+path, handler)
}

func newAPIHandler(node *core.PeernotifyNode) *apiHandler {
	h := &apiHandler{node}
	return h
}

func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Path)
	path := u.String()
	switch r.Method {
	case "GET":
		h.get(path, w, r)
	case "POST":
		h.post(path, w, r)
	case "PUT":
		h.put(path, w, r)
	case "DELETE":
		h.del(path, w, r)
	case "PATCH":
		h.patch(path, w, r)
	default:
		return
	}
}

func (h *apiHandler) GETRegister(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *apiHandler) GETVerify(w http.ResponseWriter, r *http.Request) {
	return
}
