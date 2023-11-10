package manager

import (
	"fmt"
	"sync"
)

type PDU interface {
	Get(oid []string) ([]string, error)
	Set(oid string, state int) error
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

type Manager struct {
	Devices   []Device
	NowDevice Device
	mu        sync.Mutex
	PDU
}

func NewManager(m PDU, conf Config) *Manager {
	var devices []Device
	for _, val := range conf.Devices {
		d := Device{name: val.Name, ip: val.IP, ports: []portInfo{{oid: val.Ports[0].OID, port: val.Ports[0].Port}}}
		devices = append(devices, d)
	}
	return &Manager{
		PDU: m,
		Devices: devices,
	}
}

func (m *Manager) SearchDevice(name string) (Device, error) {
	fmt.Println("New: ", m)
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
