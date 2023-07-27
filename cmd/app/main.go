package main

import (
	"flag"
	"os"
	"snmp_get/internal/api"
	"snmp_get/internal/driver"
	"snmp_get/internal/manager"
	"snmp_get/internal/server"

	"gopkg.in/yaml.v3"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "configPath", "config/Config.yml", "path")
}

func main() {
	flag.Parse()
	cfg := driver.NewConfig()
	file, err := os.Open(confPath)
	if err != nil {
		return
	}
	yaml.NewDecoder(file).Decode(&cfg)

	new := driver.New(*cfg)
	newapi := api.NewClientAPI()
	manager := manager.NewManager(new)
	api.RunAPI(manager)
	server.StartServer(newapi)

}
