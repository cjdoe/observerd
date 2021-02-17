package storage

import (
	"io"
)

func ensureClose(i io.Closer) {
	if i != nil {
		_ = i.Close()
	}
}
