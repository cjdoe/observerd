package storage

import (
	"github.com/vTCP-Foundation/observerd/common/e"
	conf "github.com/vTCP-Foundation/observerd/common/settings"
	"github.com/vTCP-Foundation/observerd/common/tests"
	"github.com/vTCP-Foundation/observerd/producer/settings"
	"testing"
)

var (
	ClickHouseInterface = conf.Network{
		Host: "172.17.0.2",
		Port: 9000,
	}
)

func EnsureEmptyStorage(t *testing.T) (h *Handler) {
	h, err := New(&settings.StorageConfig{Network: ClickHouseInterface})
	tests.InterruptIfError(err, t)

	{
		_, err = h.connection.Exec("drop table if exists blocks;")
		tests.InterruptIfError(err, t)

		err = h.ensureBlocksTable()
		tests.InterruptIfError(err, t)
	}

	{
		_, err = h.connection.Exec("drop table if exists records;")
		tests.InterruptIfError(err, t)

		_, err = h.connection.Exec("drop table if exists records_buffer;")
		tests.InterruptIfError(err, t)

		err = h.ensureRecordsTable()
		tests.InterruptIfError(err, t)
	}

	return
}

func TestLastBlockNumberWithoutBlocks(t *testing.T) {
	h := EnsureEmptyStorage(t)
	_, err := h.LastBlockNumber()
	if err != e.ErrNoBlocks {
		t.Fatal("Expected error wasn't returned")
	}
}

// todo: add test with blocks added
