package config

type Master struct {
	Version string            `yaml:"version"`
	Config  map[string]string `yaml:"config"`
	Service []string          `yaml:"services"`
}

func (m Master) HasService(service string) bool {
	for _, ele := range m.Service {
		if ele == service {
			return true
		}
	}
	return false
}
