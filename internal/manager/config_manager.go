package manager


type Config struct {
	Devices []struct {		
		IP   string              	`yaml:"ip"`
		Name string                	`yaml:"name"`
		Ports []struct {
			OID string 					`yaml:"oid"`
			Port int 					`yaml:"portnum"`
		} `yaml:"ports"`
	} `yaml:"devices"`
}

func NewConfig() *Config {
	return &Config{}
}