package main

// var (
// 	confPath string
// )

// func init() {
// 	flag.StringVar(&confPath, "configPath", "/home/manage/Documents/snmp_get/config/Config.yml", "path")
// }

// func main() {
// 	c := make(chan int)
// 	flag.Parse()
// 	cfg := snmptouch.NewConfig()
// 	file, err := os.Open(confPath)
// 	if err != nil {
// 		return
// 	}
// 	yaml.NewDecoder(file).Decode(&cfg)
// 	a := snmptouch.NewTrap(*cfg)

// 	go a.Trap(c)

// 	c <- 2

// }
