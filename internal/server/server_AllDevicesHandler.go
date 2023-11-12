package server

import (
	"net/http"
	//"fmt"
)

func (h *ManageToAPI) AllDevicesHandler(w http.ResponseWriter, r *http.Request) {
	//h.CheckAuth(w, r)
	res, err := h.Client.ClientGetAllDevices()
	if err != nil {
		return
	}
	tmplManager.Execute(w, res)
}
