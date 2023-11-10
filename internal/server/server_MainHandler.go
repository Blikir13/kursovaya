package server

import "net/http"

func (h *ManageToAPI) MainHandler(w http.ResponseWriter, r *http.Request) {
	tmplMain.Execute(w, nil)
}
