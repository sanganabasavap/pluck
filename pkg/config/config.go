package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadGlobalConfig(path string) (*Master, error) {
	readFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error while reading file %s %s\n", path, err)
		return nil, err
	}
	master := new(Master)
	err = yaml.Unmarshal(readFile, &master)
	if err != nil {
		fmt.Printf("error while reading file %s\n", err)
		return nil, err
	}
	return master, nil
}
