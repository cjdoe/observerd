package logchain

import (
	"github.com/vTCP-Foundation/observerd/common/binlog"
	"github.com/vTCP-Foundation/observerd/common/crypto"
	"math/big"
	"time"
)

type BlockHeader struct {
	Index             *big.Int
	Hash              *crypto.Hash
	RecordsHash       *crypto.Hash
	PrevBlockHash     *crypto.Hash
	Signature         *crypto.LamportSig
	NextRecordPubKey  *crypto.LamportPubKey
	DateTimeGenerated time.Time
}

type Block struct {
	Header  *BlockHeader
	Records []*binlog.Record
}
