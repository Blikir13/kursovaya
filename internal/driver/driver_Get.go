package driver

import (
	"fmt"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func (s *SNMP) Get(oid []string) ([]string, error) {
	con := &g.GoSNMP{
		Target:    s.DeviceIP,
		Port:      s.Config.Server.Port,
		Version:   s.Config.Server.Version,
		Community: "public",
		Timeout:   time.Duration((s.Config.Server.Timeout)) * time.Second,
		Retries:   s.Config.Server.Retries,
	}

	err := con.Connect()
	defer con.Conn.Close()
	if err != nil {
		fmt.Printf("Connect() %v", err)
		return nil, err
	}

	var ret []string

	result, err := con.Get(oid)
	if err != nil {
		fmt.Printf("Get() %v", err)
		return nil, err
	}

	// Forming Result
	for _, v := range result.Variables {
		switch v.Type {
		case g.OctetString:
			ret = append(ret, fmt.Sprint(v.Value.([]byte)))
		default:
			ret = append(ret, fmt.Sprint(State[g.ToBigInt(v.Value).Int64()]))
		}
	}

	return ret, nil
}
