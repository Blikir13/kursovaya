package driver

import (
	"fmt"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func (s *SNMP) Set(oid string, state int) error {

	con := &g.GoSNMP{
		Target:    s.DeviceIP,
		Port:      s.Config.Server.Port,
		Version:   s.Config.Server.Version,
		Community: "private",
		Timeout:   time.Duration(s.Config.Server.Timeout) * time.Second,
		Retries:   s.Config.Server.Retries,
	}
	err := con.Connect()
	defer con.Conn.Close()
	if err != nil {
		fmt.Printf("Connect() %v", err)
		return err
	}

	// Forming PDU
	var pdu = []g.SnmpPDU{{
		Name:  oid,
		Type:  g.Integer,
		Value: state,
	}}
	// Set()
	_, err = con.Set(pdu)
	if err != nil {
		fmt.Printf("Get() %v", err)
		return err
	}

	return nil
}
