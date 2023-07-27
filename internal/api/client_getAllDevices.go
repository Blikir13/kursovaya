package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snmp_get/internal/manager"
)

func (h *ClientAPI) ClientGetAllDevices() (manager.JsonRes, error) {
	res, err := http.Get(localhostURL + mainURL)
	if err != nil {
		fmt.Printf("error: %v", err)
		return manager.JsonRes{}, err
	}

	m, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error read body: %v", err)
		return manager.JsonRes{}, err
	}
	ret := manager.JsonRes{}
	json.Unmarshal(m, &ret)
	return ret, nil
}
