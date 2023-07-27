package server

import "net/http"

func (h *ManageToAPI) RootHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.Client.ClientGetAllDevices()
	if err != nil {
		return
	}
	tmplMain.Execute(w, res)
}
