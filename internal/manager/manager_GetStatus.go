package manager

import "fmt"

type PortsInfo struct {
	Port  string `json:"Port"`
	State string `json:"State"`
}

type GetResult struct {
	Result []PortsInfo
	Error  string `json:"Error"`
}

func (m *Manager) GetStatus(name string) (GetResult, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	device, err := m.SearchDevice(name)
	m.NowDevice = device
	m.PDU.InitIP(device.ip)
	if err != nil {
		return GetResult{}, err
	}
	var allOid []string
	res := GetResult{}
	for _, val := range device.ports {
		allOid = append(allOid, val.oid)
	}
	resGet, err := m.PDU.Get(allOid)
	if err != nil {
		return GetResult{Error: fmt.Sprint(err)}, err
	}
	for i, val := range device.ports {
		res.Result = append(res.Result, PortsInfo{Port: fmt.Sprint(val.port), State: resGet[i]})
	}
	return res, nil
}
