package config

type Master struct {
	Version string            `yaml:"version"`
	Common map[string]string `yaml:"common"`
	Cpanel map[string]string `yaml:"cpanel"`
	Gateway map[string]string `yaml:"gateway"`
	Indexer map[string]string `yaml:"indexer"`
	Search map[string]string `yaml:"search"`
	Logwatcher map[string]string `yaml:"logwatcher"`
	HistoricSearch map[string]string `yaml:"historic-search"`
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

func (m Master) GetValue(data map[string]string, key string) string {
	for k, v := range data {
		if k == key {
			return v
		}
	}
	return "null"
}
