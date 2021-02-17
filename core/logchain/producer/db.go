package producer

import (
	"fmt"
	"github.com/vTCP-Foundation/observerd/core/database"
	"github.com/vTCP-Foundation/observerd/core/ec"
	"time"
)

func (p *Producer) fetchLastBlockNumber() (number uint64, err error) {
	tx, err := database.NewTransaction()
	if err != nil {
		return
	}

	defer database.RollbackSafely(tx)

	query := "SELECT number FROM blocks ORDER BY number DESC LIMIT 1"
	rows, err := database.Query(tx, query)
	if err != nil {
		return
	}

	if !rows.Next() {
		// Blocks table has no even one entry.
		// Last block number is 0.
		return
	}

	err = rows.Scan(&number)
	return
}

func (p *Producer) fetchLastBlockTimestamp() (timestamp time.Time, err error) {
	tx, err := database.NewTransaction()
	if err != nil {
		return
	}

	defer database.RollbackSafely(tx)

	query := "SELECT timestamp FROM blocks ORDER BY number DESC LIMIT 1"
	rows, err := database.Query(tx, query)
	if err != nil {
		return
	}

	if !rows.Next() {
		err = fmt.Errorf("can't fetch last block timestamp: %w", ec.ErrNoData)
		return
	}

	err = rows.Scan(&timestamp)
	return
}
