package driver

import (
	"sync"
)

var State = []string{"Err", "On", "Off", "No installed"}

type SNMP struct {
	mu       sync.Mutex
	Config   Config
	DeviceIP string
}

func New(conf Config) *SNMP {
	return &SNMP{
		Config: conf,
	}
}

func (s *SNMP) InitIP(ip string) {
	s.DeviceIP = ip
}
