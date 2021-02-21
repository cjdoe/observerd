package log

import "github.com/vTCP-Foundation/observerd/core/p2p"

type BinLog struct {
}

func (log *BinLog) AppendClaim(claim *p2p.Claim) (err error) {
	return err
}

func (log *BinLog) AppendTLS(claim *p2p.TSL) (err error) {
	return err
}
