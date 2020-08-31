package storage

import (
	"github.com/vTCP-Foundation/observerd/core/settings"
)

func InitStorage() (driver Driver) {
	if settings.Conf.Debug {
		driver = NewLocalFiles()

	} else {
		panic("not implemented")
	}

	return
}
