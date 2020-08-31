package storage

// todo: implement for the production purposes
type Ceph struct {
}

func NewCeph() (storage *Ceph, err error) {
	return
}

func (s *Ceph) GetBlock(index uint64) (record Record, err error) {
	return
}

func (s *Ceph) AppendBlock(record Record) (index uint64, err error) {
	return
}
