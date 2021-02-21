package producer

import (
	"errors"
	"fmt"
	"github.com/vTCP-Foundation/observerd/core/ec"
	"github.com/vTCP-Foundation/observerd/core/utils"
	crypto "github.com/vTCP-Foundation/observerd/go-lamport-crypto"
	"time"
)

type Producer struct {
	nextBlockDelayFirstRun bool
}

func New() *Producer {
	return &Producer{
		nextBlockDelayFirstRun: true,
	}
}

func (p *Producer) Run() (errorsFlow <-chan error) {
	flow := make(chan error)
	go p.processInternalLoop(flow)
	return flow
}

func (p *Producer) processInternalLoop(flow chan error) {
	var err error
	defer func() {
		flow <- err
	}()

	err = p.ensureGenesisBlock()
	if err != nil {
		return
	}

	var delay time.Duration
	for {
		delay, err = p.calculateNextBlockDelay()
		if err != nil {
			// todo: add logger here
			return
		}

		time.Sleep(delay)
		err = p.createNextBlock()
		if err != nil {
			return
		}

	}
}

func (p *Producer) ensureGenesisBlock() (err error) {
	_, err = p.fetchLastBlockTimestamp()
	if errors.Is(err, ec.ErrNoData) {
		// There is no blocks yet.
		err = p.writeGenesisBlock()
		if err != nil {
			err = fmt.Errorf("can't generate genesis block: %w", err)
			return
		}
	}

	return
}

func (p *Producer) writeGenesisBlock() (err error) {
	// Genesis block contains no data.
	// So it needs some random hash as a salt for consequent hashes of all other blocks.
	h, err := crypto.GenerateRandomHash()
	if err != nil {
		err = fmt.Errorf("can't generate random hash for the genesis block: %w", err)
		return
	}

	sig, nextBlockPubKey, err := p.signBlockHash(h, 0)
	if err != nil {
		err = fmt.Errorf("can't sign hash of the genesis block: %w", err)
		return
	}

	err = p.appendBlock(0, sig, crypto.Hash{}, h, nextBlockPubKey)
	return
}

func (p *Producer) calculateNextBlockDelay() (delay time.Duration, err error) {
	// [Optimisation]
	// There is no need to check database each time when next block delay is fetched.
	// The delay could differ from standard only in case if process has been restarted,
	// and database contains last block, that has been generated less than "delay" amount of time ago.
	// In this case delay must be aligned to keep blocks generation time window the same for other observers.
	if !p.nextBlockDelayFirstRun {
		return ec.LogChainRoundDelay, nil
	}

	defer func() {
		// After first run, this method must return the constant.
		// But this mode must be enabled only in case if there was no errors.
		if err != nil {
			p.nextBlockDelayFirstRun = false
		}
	}()

	lastBlockTimestamp, err := p.fetchLastBlockTimestamp()
	if err != nil {
		if errors.Is(err, ec.ErrNoData) {
			// Blocks table contains no data (no last block timestamp is present).
			// Current block would be genesis block, so no delay must be returned.
			return 0, nil
		}

		return
	}

	now := utils.UTCNow()
	deadline := lastBlockTimestamp.Add(ec.LogChainRoundDelay)
	if deadline.Before(now) {
		return 0, nil
	}

	return deadline.Sub(now), nil
}

func (p *Producer) createNextBlock() (err error) {
	lastBlockNumber, err := p.fetchLastBlockNumber()
	if err != nil {
		return
	}

	err = p.appendBlock(lastBlockNumber+1, &crypto.LamportSig{}, crypto.Hash{}, crypto.Hash{}, &crypto.LamportPubKey{})
	return
}
