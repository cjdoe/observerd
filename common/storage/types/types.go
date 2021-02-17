package types

import (
	"bytes"
	"github.com/vTCP-Foundation/observerd/common/crypto"
	"time"
)

type Block struct {
	Version        uint16
	Number         uint64
	Hash           crypto.Hash
	PrevHash       crypto.Hash
	Signature      crypto.LamportSig
	NextSigPubKey  crypto.LamportPubKey
	RecordsNumbers []uint64
	Generated      time.Time
}

func (b *Block) Cmp(other *Block) bool {
	if b.Version != other.Version {
		return false
	}

	if b.Number != other.Number {
		return false
	}

	if bytes.Compare(b.Hash[:], other.Hash[:]) != 0 {
		return false
	}

	if bytes.Compare(b.PrevHash[:], other.PrevHash[:]) != 0 {
		return false
	}

	if bytes.Compare(b.Signature[:], other.Signature[:]) != 0 {
		return false
	}

	if bytes.Compare(b.NextSigPubKey[:], other.NextSigPubKey[:]) != 0 {
		return false
	}

	if len(b.RecordsNumbers) != len(other.RecordsNumbers) {
		return false
	}

	for i, number := range b.RecordsNumbers {
		if other.RecordsNumbers[i] != number {
			return false
		}
	}

	// todo: [bug] [clickhouse]
	//		 Clickhouse stores all values of the DateTime-column with common time zone
	//		 (stored as a column metadata option). This time zone parameter must be specified
	//		 on a column initialisation moment and influences all queries, that select data from it.
	// 		 For some reason, equal date time objects becomes different after writing/selecting from the CH,
	//		 and leads to broken validation.
	//		 This validation rule is disabled cause in case if dates are different -
	//		 hashes of the blocks would be different (and we have a rule for this check).
	//		 But, still this cae must be studied and solved.
	//
	//if b.Generated != other.Generated {
	//	return false
	//}

	return true
}

type Record struct {
	Number    uint64
	Generated time.Time
	Type      uint16
	Data      string
}
