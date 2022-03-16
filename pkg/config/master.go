package config

import (
    "strings"
    "strconv"
)

type Master struct {
	Version string            `yaml:"version"`
	Service []string          `yaml:"services"`
	Common map[string]string `yaml:"common"`
	Cpanel map[string]string `yaml:"cpanel"`
	Gateway map[string]string `yaml:"gateway"`
	Indexer map[string]string `yaml:"indexer"`
	Search map[string]string `yaml:"search"`
	Logwatcher map[string]string `yaml:"logwatcher"`
	HistoricSearch map[string]string `yaml:"historic-search"`
	Resources map[string]string `yaml:"resources"`
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

func (m Master) GetResource(data map[string]string, key string) string {
        for k, v := range data {
                if (k == key) && (strings.Contains(k,"cpu")) {
			if (v[len(v)-1:] == "m") {
				//strings.Contains(v,"m")
				num := v[:len(v)-1]
				value, _ := strconv.Atoi(num)
				cpu := strconv.FormatFloat(float64(value)/1000,'g',2,32)
				return cpu
			}
                        return v
                }
                if (k == key) && (strings.Contains(k,"memory")) {
                        if (v[len(v)-1:] == "i") {
                                mem := v[:len(v)-1]
                                return mem
                        }
                        return v
                }
        }
        return "null"
}
