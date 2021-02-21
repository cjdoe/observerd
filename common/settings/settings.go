package settings

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Conf *Settings
)

type InterfacesConfig struct {
	// Interface for vTCP nodes (p2p layer).
	Nodes Network `yaml:"nodes"`

	// Interface for other vTCP observers and public data access.
	Public Network `yaml:"public"`
}

type DatabaseConfig struct {
	Network     Network     `yaml:"network"`
	Credentials Credentials `yaml:"credentials"`
	Name        string      `yaml:"name"`
}

func (c DatabaseConfig) ConnectionCredentials() string {
	return fmt.Sprint("postgres://", c.Credentials.User, ":",
		c.Credentials.Pass, "@", c.Network.Host, ":", c.Network.Port, "/", c.Name)
}

type Settings struct {
	Interfaces InterfacesConfig `yaml:"interfaces"`
	Database   DatabaseConfig   `yaml:"database"`

	Debug bool `yaml:"debug"`
}

func Load() (err error) {
	data, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		return
	}

	return
}
