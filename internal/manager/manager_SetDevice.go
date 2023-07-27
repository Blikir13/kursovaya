package manager

func (m *Manager) SetDevice(num int, state int, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	device, err := m.SearchDevice(name)
	m.PDU.InitIP(device.ip)
	if err != nil {
		return err
	}

	m.NowDevice = device
	var oid string
	oid, err = m.SearchPortOid(num)
	if err != nil {
		return err
	}

	err = m.PDU.Set(oid, state)
	if err != nil {
		return err
	}
	return nil
}
