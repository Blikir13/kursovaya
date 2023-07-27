package driver

import "github.com/gosnmp/gosnmp"

type Config struct {
	Server struct {
		IP      string             `yaml:"ip"`
		Port    uint16             `yaml:"port"`
		Version gosnmp.SnmpVersion `yaml:"version"`
		Timeout int                `yaml:"timeout"`
		Retries int                `yaml:"retries"`
	} `yaml:"server"`
}

func NewConfig() *Config {
	return &Config{}
}
