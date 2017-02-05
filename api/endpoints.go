package api

import "net/http"

func (h *apiHandler) get(path string, w http.ResponseWriter, r *http.Request) {
	switch path {
	case "/register":
		h.GETRegister(w, r)
	case "/verify":
		h.GETVerify(w, r)
	default:
		return
	}

}

func (h *apiHandler) put(path string, w http.ResponseWriter, r *http.Request) {

}

func (h *apiHandler) post(path string, w http.ResponseWriter, r *http.Request) {

}

func (h *apiHandler) del(path string, w http.ResponseWriter, r *http.Request) {

}

func (h *apiHandler) patch(path string, w http.ResponseWriter, r *http.Request) {

}
