package api

import (
	"encoding/json"
	"io"
	"net/http"

	"snmp_get/internal/manager"
)

func (h *ClientAPI) ClientGetAllDevices() (manager.JsonRes, error) {
	res, err := http.Get(localhostURL + mainURL)
	if err != nil {
		return manager.JsonRes{}, err
	}

	m, err := io.ReadAll(res.Body)
	if err != nil {
		return manager.JsonRes{}, err
	}
	ret := manager.JsonRes{}
	json.Unmarshal(m, &ret)
	return ret, nil
}
