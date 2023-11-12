package server

import (
	"net/http"
	"fmt"
)

func (h *ManageToAPI) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("password: ", r.FormValue("login"), r.FormValue("password"))
	login := r.FormValue("login")
	password := r.FormValue("password")
	if h.DataBase.Login(login, password) != nil {
		h.LoginHandler(w, "Пароль!")
	}
	h.Auth = true
	http.Redirect(w, r, "/alldevices", http.StatusSeeOther)
}