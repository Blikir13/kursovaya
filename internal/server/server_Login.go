package server

import (
	"net/http"
)

func (h *ManageToAPI) LoginHandler(w http.ResponseWriter, message string) {
	type answer struct {
		Message string
	}
	data := answer{message}
	tmplLogin.Execute(w, data)
}