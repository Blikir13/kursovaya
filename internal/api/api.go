package api

import (
	"net/http"
	"snmp_get/internal/manager"

	"github.com/gorilla/mux"
)

const (
	mainURL      = "/api/devices"
	localhostURL = "http://localhost:8000"
)

type Manager interface {
	GetStatus(name string) (manager.GetResult, error)
	GetAllStatus() (manager.JsonRes, error)
	SetDevice(num int, state int, name string) error
	SearchDevice(name string) (manager.Device, error)
}

type APIHandler struct {
	Manager
}

func NewAPI(m Manager) *APIHandler {
	return &APIHandler{
		Manager: m,
	}
}

type ClientAPI struct {
}

func NewClientAPI() *ClientAPI {
	return &ClientAPI{}
}

func RunAPI(m Manager) {
	r := mux.NewRouter()
	h := NewAPI(m)

	r.HandleFunc(mainURL, h.GetAllDevices).Methods("GET")
	r.HandleFunc(mainURL+"/{name}", h.getDevice).Methods("GET")
	r.HandleFunc(mainURL+"/{name}/{port}&{state}", h.setDevice).Methods("POST")
	go http.ListenAndServe(":8000", r)

}
