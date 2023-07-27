package server

import (
	"fmt"
	"net/http"
	"snmp_get/internal/manager"

	"html/template"
)

var (
	tmpl     *template.Template
	tmplMain *template.Template
)

type Client interface {
	ClientGetAllDevices() (manager.JsonRes, error)
	ClientGetDevices(name string) (manager.GetResult, error)
	ClientSetDevice(port string, state string, name string) error
}

type ManageToAPI struct {
	Client
	DeviceName string
}

func NewManageToAPI(m Client) *ManageToAPI {
	return &ManageToAPI{
		Client: m,
	}
}

// Init template
func init() {
	tmpl = template.Must(template.ParseFiles("/home/manage/Documents/snmp_get/static/index.html"))
	tmplMain = template.Must(template.ParseFiles("/home/manage/Documents/snmp_get/static/main.html"))
}

func StartServer(m Client) error {
	manager := NewManageToAPI(m)
	mux := http.NewServeMux()
	mux.HandleFunc("/", manager.RootHandler)
	mux.HandleFunc("/postform", manager.PostHandler)
	mux.HandleFunc("/device", manager.DeviceHandler)
	fmt.Println("StartServer()")
	err := http.ListenAndServe(":8001", mux)

	if err != nil {
		fmt.Printf("StartServer() error: %v\n", err)
		return err
	}
	return nil
}
