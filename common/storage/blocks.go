package storage

import (
	"github.com/vTCP-Foundation/observerd/common/e"
	"github.com/vTCP-Foundation/observerd/common/storage/types"
	"time"
)

func (h *Handler) WriteBlock(b *types.Block) (err error) {
	tx, err := h.connection.Begin()
	if err != nil {
		return
	}

	query := "insert into blocks (" +
		"version, number, hash, prevHash, signature, nextSigPubKey, recordsNumbers, generated" +
		") " +
		"values (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := tx.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(
		b.Version,
		b.Number,
		string(b.Hash[:]),
		string(b.PrevHash[:]),
		string(b.Signature[:]),
		string(b.NextSigPubKey[:]),
		b.RecordsNumbers,
		b.Generated.UTC())

	err = tx.Commit()
	defer ensureClose(stmt)
	return
}

func (h *Handler) FetchBlock(number uint64) (b *types.Block, err error) {
	rows, err = h.connection.Query(
		"SELECT version, number, generated, hash, prevHash, signature, nextSigPubKey, recordsNumbers "+
			"FROM blocks "+
			"WHERE number = ?  LIMIT 1", number)
	defer ensureClose(rows)
	if err != nil {
		return
	}

	blockRecordIsPresent := rows.Next()
	if !blockRecordIsPresent {
		err = e.ErrNoBlocks
	}

	b = &types.Block{}
	var hash, prevHash, signature, nextSigPbKey string
	err = rows.Scan(&b.Version, &b.Number, &b.Generated, &hash, &prevHash, &signature, &nextSigPbKey, &b.RecordsNumbers)
	if err != nil {
		return
	}

	b.Generated = b.Generated.UTC()

	copy(b.Hash[:], []byte(hash)[:])
	copy(b.PrevHash[:], []byte(prevHash)[:])
	copy(b.Signature[:], []byte(signature)[:])
	copy(b.NextSigPubKey[:], []byte(nextSigPbKey)[:])
	return
}

// DelaySinceLastBlock returns difference of 2 moments of time:
// UTC Now() - Last Block generation time.
// This time specifies how long the next block has no been produced.
func (h *Handler) DelaySinceLastBlock() (delay time.Duration, err error) {
	blocksCount, err := h.BlocksCount()
	if err != nil {
		return
	}

	if blocksCount == 0 {
		err = e.ErrNoBlocks
		return
	}

	block, err := h.FetchBlock(blocksCount - 1)
	if err != nil {
		return
	}

	// Note:
	// The comparison is done using UTCNow as a basis.
	// This is needed to mitigate various potential bugs related to time-zones processing.
	delay = time.Now().UTC().Sub(block.Generated)
	return
}
