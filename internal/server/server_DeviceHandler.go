package server

import (
	"fmt"
	"net/http"
)

func (h *ManageToAPI) DeviceHandler(w http.ResponseWriter, r *http.Request) {
	h.DeviceName = r.FormValue("device")
	res, err := h.Client.ClientGetDevices(h.DeviceName)
	if err != nil {
		fmt.Printf("error FormResult(): %v", err)
		return
	}
	tmpl.Execute(w, res)
}
