package settings

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Conf *Settings
)

type Interfaces struct {
	Nodes  networkInterface `yaml:"nodes"`
	Public networkInterface `yaml:"public"`
}

type Settings struct {
	Debug      bool       `yaml:"debug"`
	Interfaces Interfaces `yaml:"interfaces"`
}

func Load() (err error) {
	Conf = &Settings{}

	data, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		return
	}

	return
}
