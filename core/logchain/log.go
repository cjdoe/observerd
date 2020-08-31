package logchain

import (
	"github.com/vTCP-Foundation/observerd/core/storage"
	"time"
)

type Log struct {
	storage storage.Driver
}

func New() (log *Log, err error) {
	log = &Log{
		storage: storage.InitStorage(),
	}

	err = log.storage.RestoreOrInit()
	if err != nil {
		return
	}

	height, err := log.storage.Height()
	if err != nil {
		return
	}

	if height == 0 {

	}

	return
}

func (log *Log) writeGenesis() (err error) {
	header := Header{
		Index:             0,
		Hash:              nil,
		PrevBlockHash:     nil,
		Signature:         nil,
		NextRecordPubKey:  nil,
		DateTimeGenerated: time.Time{},
	}

	block := Block{
		Header:  nil,
		Records: nil,
	}
}
