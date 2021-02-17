package settings

import "fmt"

type Network struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

func (ni Network) Interface() string {
	return fmt.Sprint(ni.Host, ":", ni.Port)
}

type Credentials struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}
