package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *APIHandler) setDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	state, _ := strconv.Atoi(params["state"])
	port, _ := strconv.Atoi(params["port"])
	if err := h.Manager.SetDevice(port, state, params["name"]); err != nil {
		fmt.Printf("Error when Set(): %v\n", err)
		return
	}
}
