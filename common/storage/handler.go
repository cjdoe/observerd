package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/vTCP-Foundation/observerd/common/storage/types"
	"github.com/vTCP-Foundation/observerd/producer/settings"
)

var (
	rows *sql.Rows
)

type Handler struct {
	connection *sql.DB
}

func New(conf *settings.StorageConfig) (handler *Handler, err error) {
	dataSource := fmt.Sprint("tcp://", conf.Network.Interface())
	connection, err := sql.Open("clickhouse", dataSource)
	if err != nil {
		return
	}

	err = connection.Ping()
	if err != nil {
		return
	}

	handler = &Handler{
		connection: connection,
	}
	return
}

func (h *Handler) FetchRecords(offset uint64) (records []*types.Record, err error) {
	rows, err = h.connection.Query("SELECT number, generated, type, data FROM records WHERE number >= ?", offset)
	defer ensureClose(rows)
	if err != nil {
		return
	}

	records = make([]*types.Record, 0, 2048)
	for rows.Next() {
		r := &types.Record{}
		err = rows.Scan(&r.Number, &r.Generated, &r.Type, &r.Data)
		if err != nil {
			return
		}

		records = append(records, r)
	}

	return
}
