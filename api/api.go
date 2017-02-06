package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/golang/protobuf/jsonpb"
	"github.com/yurizhykin/peernotify/core"
	"github.com/yurizhykin/peernotify/pb"
)

const (
	API_PREFIX  = "" // "/api/"
	API_VERSION = "" // "v1"
)

type apiHandler struct {
	mux  map[string]func(http.ResponseWriter, *http.Request)
	node *core.PeernotifyNode
}

func NewAPIServer(node *core.PeernotifyNode, addr string) http.Server {
	h := newAPIHandler(node)
	h.register(map[string]func(http.ResponseWriter, *http.Request){
		"/register": h.handleRegister,
		"/verify":   h.handleVerify,
		"/forward":  h.handleForward,
	})
	return http.Server{Addr: addr, Handler: h}
}

func (h *apiHandler) register(handlers map[string]func(http.ResponseWriter, *http.Request)) {
	for path, handler := range handlers {
		h.mux[API_PREFIX+API_VERSION+path] = handler
	}
}

func newAPIHandler(node *core.PeernotifyNode) *apiHandler {
	return &apiHandler{make(map[string]func(http.ResponseWriter, *http.Request)), node}
}

func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		http.Error(w, "Malformed URL", 400)
	}
	path := u.String()
	if handler, ok := h.mux[path]; !ok {
		apiNotFound(w, r)
	} else {
		handler(w, r)
	}
}

func (h *apiHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		apiWrongMethod(w, r)
		return
	}
	var contact pb.Contact
	if err := jsonpb.Unmarshal(r.Body, &contact); err != nil {
		apiWrongBody(w, r)
		return
	}
	if err := h.node.Register(contact); err != nil {
		apiInternalError(w, r)
		return
	}
}

func (h *apiHandler) handleVerify(w http.ResponseWriter, r *http.Request) {
	apiNotFound(w, r)
}

func (h *apiHandler) handleForward(w http.ResponseWriter, r *http.Request) {
	apiNotFound(w, r)
}

// HTTP error responses
func apiNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func apiWrongMethod(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Unable to %s", r.Method), 405)
}

func apiWrongBody(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Malformed body"), 400)
}

func apiInternalError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Internal error"), 500)
}
