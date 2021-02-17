package binlog

import (
	"encoding"
	"time"
)

type Record interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	Created() time.Time
}

type BaseRecord struct {
	created time.Time
}

func NewBaseRecord() (record *BaseRecord) {
	record = &BaseRecord{created: time.Now()}
	return
}

func (record *BaseRecord) Created() time.Time {
	return record.created
}

func (record *BaseRecord) MarshalBinary() (data []byte, err error) {
	return record.created.MarshalBinary()
}

func (record *BaseRecord) UnmarshalBinary(data []byte) (err error) {
	return record.created.UnmarshalBinary(data)
}
