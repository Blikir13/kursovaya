package server

import (
	"fmt"
	"net/http"
	"html/template"

	"snmp_get/internal/manager"
)

var (
	tmpl     *template.Template
	tmplManager, tmplMain, tmplLogin *template.Template
)

type Client interface {
	ClientGetAllDevices() (manager.JsonRes, error)
	ClientGetDevices(name string) (manager.GetResult, error)
	ClientSetDevice(port string, state string, name string) error
}

type DataBase interface {
	Login (login, Password string) error
}

type ManageToAPI struct {
	Client
	DataBase
	DeviceName string
	Auth bool
}

func NewManageToAPI(m Client, d DataBase) *ManageToAPI {
	return &ManageToAPI{
		Client: m,
		DataBase: d,
		Auth: false,
	}
}

// Init template
func init() {
	tmpl = template.Must(template.ParseFiles("/Users/mac/Desktop/practice/static/index.html"))
	tmplManager = template.Must(template.ParseFiles("/Users/mac/Desktop/practice/static/manager.html"))
	tmplMain = template.Must(template.ParseFiles("/Users/mac/Desktop/practice/static/main.html"))
	tmplLogin = template.Must(template.ParseFiles("/Users/mac/Desktop/practice/static/login.html"))
}

func StartServer(m Client, d DataBase) error {
	manager := NewManageToAPI(m, d)
	mux := http.NewServeMux()
	mux.HandleFunc("/", manager.MainHandler)
	mux.HandleFunc("/alldevices", manager.AllDevicesHandler)
	mux.HandleFunc("/postform", manager.PostHandler)
	mux.HandleFunc("/device", manager.DeviceHandler)
	mux.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		manager.LoginHandler(rw, "")
	})
	mux.HandleFunc("/postlogin", manager.PostLoginHandler)
	fmt.Println("StartServer()")
	err := http.ListenAndServe(":8001", mux)

	if err != nil {
		fmt.Printf("StartServer() error: %v\n", err)
		return err
	}
	return nil
}

func (h *ManageToAPI) CheckAuth(w http.ResponseWriter, r *http.Request) {
	if h.Auth == false {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
