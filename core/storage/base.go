package storage

type Record interface {
	Data() []byte
}

type Driver interface {
	RestoreOrInit() (err error)
	Height() (height uint64, err error)
	GetBlock(index uint64) (record Record, err error)
	AppendBlock(record Record) (index uint64, err error)
}
