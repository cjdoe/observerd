package storage

// todo: implement for the tests purposes
type LocalFiles struct {
}

func NewLocalFiles() (storage *LocalFiles) {
	return
}

func (driver *LocalFiles) RestoreOrInit() (err error) {
	return
}

func (driver *LocalFiles) Height() (height uint64, err error) {
	// todo: implement me back
	height = 0
	return
}

func (driver *LocalFiles) GetBlock(index uint64) (record Record, err error) {
	return
}

func (driver *LocalFiles) AppendBlock(record Record) (index uint64, err error) {
	return
}
