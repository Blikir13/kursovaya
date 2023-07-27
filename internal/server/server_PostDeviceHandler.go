package server

import (
	"net/http"
)

func (h *ManageToAPI) PostHandler(w http.ResponseWriter, r *http.Request) {
	h.Client.ClientSetDevice(r.FormValue("portnum"), r.FormValue("state"), h.DeviceName)
	http.Redirect(w, r, "/device?device="+h.DeviceName, http.StatusSeeOther)
}
