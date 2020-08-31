package settings

import "fmt"

type networkInterface struct {
	Address string `yaml:"address"`
	Port    uint16 `yaml:"port"`
}

func (ni networkInterface) Interface() string {
	return fmt.Sprint(ni.Address, ":", ni.Port)
}
