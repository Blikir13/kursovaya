package main

import (
	"flag"
	"os"
	"fmt"

	"gopkg.in/yaml.v3"

	"snmp_get/internal/api"
	"snmp_get/internal/driver"
	"snmp_get/internal/manager"
	"snmp_get/internal/server"
	"snmp_get/internal/db"
)

var (
	confPath string
	managerConfPath string
)

func init() {
	flag.StringVar(&confPath, "configPath", "config/Config.yml", "path")
	flag.StringVar(&managerConfPath, "configPath1", "config/Config_manager.yml", "path")
}

func main() {
	flag.Parse()
	cfg := driver.NewConfig()
	manager_cfg := manager.NewConfig()
	file, err := os.Open(confPath)
	if err != nil {
		return
	}
	ManagerConfFile, err := os.Open(managerConfPath)
	if err != nil {
		return
	}
	yaml.NewDecoder(ManagerConfFile).Decode(&manager_cfg)
	yaml.NewDecoder(file).Decode(&cfg)


	base := db.NewDB()
	fmt.Println( base.Login("12", "12"))

	new := driver.New(*cfg)
	newapi := api.NewClientAPI()
	manager := manager.NewManager(new, *manager_cfg)
	//manager.Monitoring()
	api.RunAPI(manager)
	server.StartServer(newapi, base)

}
