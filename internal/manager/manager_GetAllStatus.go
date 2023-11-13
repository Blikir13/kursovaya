package manager

import "fmt"

type ManagerOutlet struct {
	Names  string `json:"Name"`
	IPs    string `json:"Ip"`
	States string `json:"States"`
	Error  string `json:"Error"`
}

type JsonRes struct {
	Out []ManagerOutlet
}

func (m *Manager) GetAllStatus() (JsonRes, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	result := JsonRes{}
	for _, val := range m.Devices {
		res, err := m.PDU.CheckState(val.ip)
		if err != nil {
			return JsonRes{}, err
		}
		if res {
			result.Out = append(result.Out, ManagerOutlet{Names: val.name, IPs: val.ip, States: "Подключен", Error: ""})
		} else {
			result.Out = append(result.Out, ManagerOutlet{Names: val.name, IPs: val.ip, States: "Не подключен", Error: fmt.Sprintf("%v", err)})
		}
	}
	return result, nil
}
