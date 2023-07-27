package driver

import (
	"fmt"
	"time"

	g "github.com/gosnmp/gosnmp"
)

// возвращает состояние всех устройств
func (s *SNMP) CheckState(ip string) (bool, error) {

	con := &g.GoSNMP{
		Target:    ip,
		Port:      s.Config.Server.Port,
		Version:   s.Config.Server.Version,
		Community: "private",
		Timeout:   time.Duration(s.Config.Server.Timeout) * time.Second,
	}

	err := con.Connect()
	defer con.Conn.Close()

	if err != nil {
		fmt.Printf("Connect() %v", err)
		return false, err
	}

	_, err = con.Get([]string{"1.3.6.1.2.1.1.5"})

	if err != nil {
		return false, nil
	}

	return true, nil
}
