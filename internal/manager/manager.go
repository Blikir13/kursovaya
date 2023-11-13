package manager

import (
	"fmt"
	"sync"
)

type PDU interface {
	Get(oid []string) ([]string, error)
	Set(oid int, state int) error
	CheckState(ip string) (bool, error)
	InitIP(ip string)
}

type portInfo struct {
	oid  string
	port int
}

type Device struct {
	name  string
	ip    string
	ports []portInfo
}

type DataBase interface {
	Write (name string, port int, port_state string, bool_change bool) error
}

type Manager struct {
	Devices   []Device
	NowDevice Device
	mu        sync.Mutex
	PDU
	DataBase
}

func NewManager(m PDU, conf Config, d DataBase) *Manager {
	var devices []Device
	var p []portInfo
	for _, val := range conf.Devices {
		for _, porti := range val.Ports{
			p = append(p, portInfo{oid: porti.OID, port: porti.Port})
		}
		d := Device{name: val.Name, ip: val.IP, ports: p}
		devices = append(devices, d)
		p = []portInfo{}
	}
	return &Manager{
		PDU: m,
		Devices: devices,
		DataBase: d,
	}
}

func (m *Manager) SearchDevice(name string) (Device, error) {
	//fmt.Println("New: ", m)
	for _, val := range m.Devices {
		if val.name == name {
			return val, nil
		}
	}
	return Device{}, fmt.Errorf("no such device: %s", name)
}

func (m *Manager) SearchPortOid(num int) (string, error) {
	for _, val := range m.NowDevice.ports {
		if val.port == num {
			return val.oid, nil
		}
	}
	return "", fmt.Errorf("no such port: %d", num)
}
