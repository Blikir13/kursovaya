package server

import (
	"net/http"
)

func (h *ManageToAPI) LoginHandler(w http.ResponseWriter, r *http.Request,message string) {
	type answer struct {
		Message string
	}
	data := answer{message}
	if h.Auth == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tmplLogin.Execute(w, data)
}