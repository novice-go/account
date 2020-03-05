package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func InitConfig(file string, resp interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, resp); err != nil {
		return err
	}

	return nil
}

