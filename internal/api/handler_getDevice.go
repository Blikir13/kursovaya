package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Output struct {
	Range []string `json:"Range"`
	Ports []int    `json:"Ports"`
	Error string   `json:"Error"`
}

func (h *APIHandler) getDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	out, err := h.Manager.GetStatus(params["name"])
	if err != nil {
		fmt.Printf("error FormResult(): %v", err)
		return
	}
	json.NewEncoder(w).Encode(out)
}
