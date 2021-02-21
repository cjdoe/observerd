package producer

import (
	"encoding/hex"
	"fmt"
	"github.com/vTCP-Foundation/observerd/core/database"
	"github.com/vTCP-Foundation/observerd/core/ec"
	crypto "github.com/vTCP-Foundation/observerd/go-lamport-crypto"
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

func (p *Producer) appendBlock(
	number uint64, sig *crypto.LamportSig, prevBlockHash, hash crypto.Hash,
	nextPubKey *crypto.LamportPubKey) (err error) {

	tx, err := database.NewTransaction()
	defer database.RollbackSafely(tx)
	if err != nil {
		return
	}

	blockCreationQuery := fmt.Sprint(
		"INSERT INTO blocks (number, prev_block_hash, hash, sig, next_block_pub_key) ",
		"VALUES (",
		number, ", ",
		encodeHexField(prevBlockHash[:]), ", ",
		encodeHexField(hash[:]), ", ",
		encodeHexField(sig[:]), ", ",
		encodeHexField(nextPubKey[:]),
		")")

	err = database.Exec(tx, blockCreationQuery)
	if err != nil {
		return
	}

	err = database.Commit(tx)
	return
}

func encodeHexField(bytes []byte) (sql string) {
	hexRepresentation := hex.EncodeToString(bytes)
	return fmt.Sprint("decode('", hexRepresentation, "', 'hex') ")
}
