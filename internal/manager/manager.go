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

func NewManager(m PDU) *Manager {
	return &Manager{
		PDU: m,
		Devices: []Device{{name: "apc1", ip: "40.1.0.73", ports: []portInfo{{oid: ".1.3.6.1.4.1.318.1.1.10.3.3.1.1.3.1", port: 1}}},
			{name: "apc2", ip: "40.1.0.55", ports: []portInfo{{oid: ".1.3.6.1.4.1.318.1.1.4.4.2.1.3.23", port: 23}, {oid: ".1.3.6.1.4.1.318.1.1.4.4.2.1.3.24", port: 24}}}},
	}
}

func (m *Manager) SearchDevice(name string) (Device, error) {
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
