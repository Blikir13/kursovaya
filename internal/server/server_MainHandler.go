package server

import "net/http"

func (h *ManageToAPI) MainHandler(w http.ResponseWriter, r *http.Request) {
	type IfAuth struct {
		Au bool
	}
	data := IfAuth{h.Auth}
	tmplMain.Execute(w, data)
}
