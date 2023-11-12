package manager

import(
	//"fmt"
	"time"
	"strconv"
)

type keepValue struct {
	name 	   string
	port       int
	state      string
}

type Listener struct {
	Values  []keepValue
}

func NewListener() *Listener {
	return &Listener{
	}
}

func (m *Manager) Monitoring() {
	keep := NewListener()
	i := 1
	for {
		for _, k := range m.Devices {
			res, _ := m.GetStatus(k.name)
			for _, val := range res.Result {
				if i == 1 {
					a, _ := strconv.Atoi(val.Port)
					keep.Values = append(keep.Values, keepValue{name: k.name, port: a, state: val.State})
					m.DataBase.Write(k.name, a, val.State, false)
				} else {
					a, _ := strconv.Atoi(val.Port)
					if  keep.CheckChange(k.name, a, val.State) == true {
						m.DataBase.Write(k.name, a, val.State, true)
					}
				}
			}
		}
		i = 0
		time.Sleep(10 * time.Second)
	}
	
}

func (l *Listener) CheckChange(name string, port int, state string) bool {
	for i, val := range l.Values {
		if (val.name == name) && (val.port == port) {
			if val.state != state {
				l.Values[i].state = state
				return true
			}
		}
	}
	return false
}