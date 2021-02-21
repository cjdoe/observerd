package producer

import (
	"errors"
	"fmt"
	crypto "github.com/vTCP-Foundation/observerd/go-lamport-crypto"
)

const (
	KeysPoolFilePath = "keys/blocks/"
)

func (p *Producer) signBlockHash(h crypto.Hash, blockNumber int64) (
	sig *crypto.LamportSig, nextBlockPubKey *crypto.LamportPubKey, err error) {

	currentBlockOperation := fmt.Sprint("block-", blockNumber)
	nextBlockBlockOperation := fmt.Sprint("block-", blockNumber+1)

	currentBlockEngine, err := crypto.NewLamport(currentBlockOperation, KeysPoolFilePath)
	if err != nil {
		return
	}

	if blockNumber == 0 {
		err = currentBlockEngine.GenerateKeypair()
		if err != nil {
			// In case if such keypair is already present - ignore the err.
			if !errors.Is(err, crypto.ErrKeypairAlreadyPresent) {
				return
			}
		}
	}

	err = currentBlockEngine.Sign(h)
	if err != nil {
		return
	}

	sig, err = currentBlockEngine.LoadSignature()
	if err != nil {
		return
	}

	nextBlockEngine, err := crypto.NewLamport(nextBlockBlockOperation, KeysPoolFilePath)
	if err != nil {
		return
	}

	err = nextBlockEngine.GenerateKeypair()
	if err != nil {
		// In case if such keypair is already present - ignore the err.
		if !errors.Is(err, crypto.ErrKeypairAlreadyPresent) {
			return
		}
	}

	nextBlockPubKey, err = nextBlockEngine.LoadPubKey()
	return
}
