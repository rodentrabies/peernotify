package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	r "github.com/julienschmidt/httprouter"
	"github.com/yurizhykin/peernotify/core"
	"github.com/yurizhykin/peernotify/pb"
)

type apiHandler struct {
	*r.Router
	node *core.PeernotifyNode
}

func NewAPIServer(node *core.PeernotifyNode, addr string) http.Server {
	h := &apiHandler{r.New(), node}

	// Register API handlers
	h.POST("/register", h.Register)
	h.GET("/verify/:id", h.Verify)
	h.POST("/forward", h.Forward)

	return http.Server{Addr: addr, Handler: h}
}

//------------------------------------------------------------------------------
// API handlers
func (h *apiHandler) Register(w http.ResponseWriter, r *http.Request, _ r.Params) {
	// Decode
	var contact pb.Contact
	if err := jsonpb.Unmarshal(r.Body, &contact); err != nil {
		apiWrongBody(w, r)
		return
	}
	// Get server URL to expect verification at
	url := "http://" + r.Host + "/verify/"
	// Run registration process
	if err := h.node.Register(contact, url); err != nil {
		apiInternalError(w, r)
		return
	}
}

func (h *apiHandler) Verify(w http.ResponseWriter, r *http.Request, ps r.Params) {
	if err := h.node.Verify(ps.ByName("id")); err != nil {
		apiInternalError(w, r)
	}
}

func (h *apiHandler) Forward(w http.ResponseWriter, r *http.Request, _ r.Params) {
	var message pb.Message
	if err := jsonpb.Unmarshal(r.Body, &message); err != nil {
		apiWrongBody(w, r)
		return
	}
	if err := h.node.Forward(message); err != nil {
		log.Println(err.Error())
		apiInternalError(w, r)
		return
	}

}

//------------------------------------------------------------------------------
// HTTP error responses
func apiWrongBody(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Malformed body"), 400)
}

func apiInternalError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Internal error"), 500)
}
