package logchain

import (
	"github.com/vTCP-Foundation/observerd/core/crypto"
	"math/big"
	"time"
)

type Record interface {
	Created() time.Time
	Data() (data []byte, err error)
}

type Header struct {
	Index             big.Int
	Hash              crypto.Hash
	PrevBlockHash     crypto.Hash
	Signature         crypto.Signature
	NextRecordPubKey  crypto.PublicKey
	DateTimeGenerated time.Time
}

type Block struct {
	Header  *Header
	Records []*Record
}
