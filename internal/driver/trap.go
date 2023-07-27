package driver

// var State = []string{"Err", "On", "Off", "No installed"}

// type keepValue struct {
// 	deviceName string
// 	oidName    string
// 	value      string
// }

// type Listener struct {
// 	Devices []device
// 	Values  []keepValue
// 	Config  Config
// }

// func NewTrap(conf Config) *Listener {
// 	return &Listener{
// 		Config: conf,
// 	}
// }

// func (l *Listener) Trap(c chan int) {
// 	var keepOID []string
// 	i := true
// 	for {
// 		for _, val := range l.Devices {
// 			for _, oid := range val.portInfo {
// 				keepOID = append(keepOID, oid.oid)        // делаем string -> []string, т к get(oid []string)
// 				resGet, err := l.GetTrap(val.ip, keepOID) // get()
// 				if err != nil {
// 					fmt.Printf("Device %s, with oid: %s unreacheable with error: %v", val.name, oid.oid, err)
// 				}
// 				if i {
// 					// на первой итерации заполняем массив начальными значениями
// 					l.Values = append(l.Values, keepValue{deviceName: val.name, oidName: oid.oid, value: resGet[0]})
// 				} else {
// 					// на последующих итерациях проверяем каждое значение на изменение
// 					if l.searchRecord(val.name, oid.oid, resGet[0]) == 0 {
// 						// какое-то действие при изменении статуса
// 						err = writeFile("change!")
// 						if err != nil {
// 							fmt.Printf("error filewrite: %v", err)
// 							return
// 						}
// 						fmt.Println("changed!")
// 					}
// 				}
// 				str := val.name + " " + oid.oid + " " + resGet[0]
// 				err = writeFile(str)
// 				if err != nil {
// 					fmt.Printf("error filewrite: %v", err)
// 					return
// 				}

// 				keepOID = nil
// 			}
// 		}

// 		i = false

// 		fmt.Println(l.Values)
// 		time.Sleep(time.Second * time.Duration(5))
// 	}
// 	<-c
// }

// func checkFileStatus(name string) (bool, error) {
// 	_, err := os.Stat(name)
// 	if err == nil {
// 		return true, nil
// 	}
// 	if errors.Is(err, os.ErrNotExist) {
// 		return false, nil
// 	}
// 	return false, err
// }

// func writeFile(s string) error {
// 	var file *os.File
// 	const name = "log.txt"
// 	check, err := checkFileStatus(name)
// 	if err != nil {
// 		return fmt.Errorf("error with checkFileStatus: %v", err)
// 	}
// 	if !check {
// 		file, err = os.Create("log.txt")
// 	} else {
// 		file, err = os.Open("/home/manage/Documents/snmp_get/log.txt")
// 	}
// 	defer file.Close()
// 	if err != nil {
// 		fmt.Printf("error create file %v", err)
// 	}
// 	time := time.Now().Format("2006-01-02 15:04:05")
// 	file.WriteString(time + " " + s + "/n")
// 	fmt.Println("print: ", s)
// 	return nil
// }

// // поиск нужного значения (по имени устройства и oid) для проверки
// func (l *Listener) searchRecord(name string, oid string, resGet string) int {
// 	for i, val := range l.Values {
// 		if (val.deviceName == name) && (val.oidName == oid) {
// 			if val.value == resGet {
// 				return 1
// 			}
// 			l.Values[i].value = resGet
// 			return 0
// 		}
// 	}
// 	return 0
// }

// // get()
// func (l *Listener) GetTrap(ip string, oid []string) ([]string, error) {
// 	con := &g.GoSNMP{
// 		Target:    ip,
// 		Port:      l.Config.Server.Port,
// 		Version:   l.Config.Server.Version,
// 		Community: "public",
// 		Timeout:   time.Duration((l.Config.Server.Timeout)) * time.Second,
// 		Retries:   l.Config.Server.Retries,
// 	}

// 	err := con.Connect()
// 	defer con.Conn.Close()
// 	if err != nil {
// 		fmt.Printf("Connect() %v", err)
// 		return nil, err
// 	}

// 	result, err := con.Get(oid)
// 	if err != nil {
// 		fmt.Printf("Get() %v", err)
// 		return nil, err
// 	}

// 	var ret []string
// 	for _, v := range result.Variables {
// 		switch v.Type {
// 		case g.OctetString:
// 			ret = append(ret, string(v.Value.([]byte)))
// 		default:
// 			ret = append(ret, State[g.ToBigInt(v.Value).Int64()])
// 		}
// 	}
// 	return ret, nil
// }
