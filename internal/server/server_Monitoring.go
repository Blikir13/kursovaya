package server

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "net/http"
	"net/http"

	// "snmp_get/internal/manager"
)

type Table struct {
	Id string `json:"Id"`
	Name  string `json:"Name"`
	State string `json:"State"`
	Port int `json:"Port"`
	Change bool `json:"Change"`
	Date string `json:"Date"`
}

type GetResult struct {
	Rows []Table
}

func (h *ManageToAPI) MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	//h.CheckAuth(w, r)
	rows := h.DataBase.GetTable()
	ret := GetResult{}
	for rows.Next() {
		p := Table{}
        rows.Scan(&p.Id, &p.Name, &p.Port, &p.State, &p.Change, &p.Date )
		ret.Rows = append(ret.Rows, p)
	}
	tmplMonitoring.Execute(w, ret)
}