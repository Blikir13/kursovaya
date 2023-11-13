package driver




type PDU interface {
	Get(oid []string) ([]string, error)
	Set(oid string, state int) error
	CheckState(ip string) (bool, error)
	InitIP(ip string)
}

type Palka struct {
	Config   Config
	DeviceIP string
	state []string
}

func NewPalka(conf Config) *Palka {
	return &Palka{
		Config: conf,
	}
}

func (s *Palka) InitIP(ip string) {
	s.DeviceIP = ip
}

func (s *Palka) Get(oid []string) ([]string, error) {
	if len(s.state) == 0{
		for i :=0; i < len(oid); i++ {
			s.state = append(s.state, "Включен")
		}
	}
	return s.state, nil
}

func (s *Palka) Set(num int, state int) error  {
	if state==1{
		s.state[num-1] = "Выключен"
	}else {
		s.state[num-1] = "Включен"
	}
	return nil
}

func (s *Palka) CheckState(ip string) (bool, error) {
	return true, nil
}
