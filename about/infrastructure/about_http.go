package infrastructure

import (
	"encoding/json"
	"net/http"
)

type AboutHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p, err := h.aboutService.Find()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&p)
}
