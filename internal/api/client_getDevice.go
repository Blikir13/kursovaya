package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snmp_get/internal/manager"
)

func (h *ClientAPI) ClientGetDevices(name string) (manager.GetResult, error) {
	res, err := http.Get(localhostURL + mainURL + "/" + name)
	if err != nil {
		fmt.Printf("error: %v", err)
		return manager.GetResult{}, err
	}

	m, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error read body: %v", err)
		return manager.GetResult{}, err
	}
	ret := manager.GetResult{}
	json.Unmarshal(m, &ret)
	return ret, nil
}
