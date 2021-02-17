package storage

import (
	"crypto/rand"
	"github.com/vTCP-Foundation/observerd/common/crypto"
	"github.com/vTCP-Foundation/observerd/common/storage/types"
	"github.com/vTCP-Foundation/observerd/common/tests"
	"testing"
	"time"
)

func randomHash() (h crypto.Hash) {
	_, _ = rand.Read(h[:])
	return
}

func randomSig() (h crypto.LamportSig) {
	_, _ = rand.Read(h[:])
	return
}

func randomPubKey() (h crypto.LamportPubKey) {
	_, _ = rand.Read(h[:])
	return
}

// This test creates random blocks with random data, writes them to the CH,
// reads them back and compares with original data structure which has been written.
// In case if originally created and selected blocks differs in some field - test reports failure.
//
// WARN: This test does not check blocks integrity
// (e.g. there is no check of hashes of the blocks form common aligned chain).
func TestBlockWriting(t *testing.T) {
	const blocksCount = 10
	blocks := make([]*types.Block, 0, blocksCount)
	for i := 0; i < blocksCount; i++ {
		blocks = append(blocks, &types.Block{
			Number:         0,
			Hash:           randomHash(),
			PrevHash:       randomHash(),
			Signature:      randomSig(),
			NextSigPubKey:  randomPubKey(),
			RecordsNumbers: []uint64{0},
			Generated:      time.Now(),
		})
	}

	for _, block := range blocks {
		h := EnsureEmptyStorage(t)
		err := h.WriteBlock(block)
		tests.InterruptIfError(err, t)

		fetchedBlock, err := h.FetchBlock(0)
		tests.InterruptIfError(err, t)

		if !block.Cmp(fetchedBlock) {
			t.Fatal("Blocks are different")
		}
	}
}
