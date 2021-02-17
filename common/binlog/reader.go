package binlog

import "os"

type Reader struct {
	file *os.File
}

func NewReader(filePath string) (reader *Reader, err error) {
	reader = &Reader{}
	reader.file, err = os.Open(filePath)
	return
}

func (reader *Reader) NextRecord() (record Record, err error) {
	panic("not implemented")
}

//var (
//	I      = &big.Int{}
//	One, _ = I.SetString("1", 10)
//)
//
//type Write struct {
//	currentBlock *BlockHeader
//	storage storage.Driver
//}
//
//func New() (log *Write, err error) {
//	log = &Write{
//		storage: storage.InitStorage(),
//	}
//
//	err = log.storage.RestoreOrInit()
//	if err != nil {
//		return
//	}
//
//	height, err := log.storage.Height()
//	if err != nil {
//		return
//	}
//
//	if height == 0 {
//
//	}
//
//	return
//}
//
//func (log *Write) FlushPool(pool *Write) (err error) {
//	nextBlockHeader := BlockHeader{
//		Index: log.currentBlock.Index,
//	}
//
//
//}
//
//func (log *Write) nextBlockHeader() (header *BlockHeader, err error) {
//	nextIndex := I.Add(log.currentBlock.Index, One)
//
//}

//func (log *Write) writeGenesis() (err error) {
//	header := BlockHeader{
//		Index:             0,
//		Hash:              nil,
//		PrevBlockHash:     nil,
//		Signature:         nil,
//		NextRecordPubKey:  nil,
//		DateTimeGenerated: time.Time{},
//	}
//
//	block := Block{
//		BlockHeader:  nil,
//		Records: nil,
//	}
//}
