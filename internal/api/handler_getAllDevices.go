package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *APIHandler) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	out, err := h.Manager.GetAllStatus()
	if err != nil {
		fmt.Printf("Error with CheckState(): %v\n", err)
		return
	}
	json.NewEncoder(w).Encode(out)

}
