package binlog

import (
	"encoding/binary"
	"errors"
	"github.com/vTCP-Foundation/observerd/common/e"
	"io"
	"os"
	"strings"
)

var (
	ErrPartialWriteOccurred = errors.New("the record has been written partially")
	ErrCorruptedLogFile     = errors.New("log file contains data, but it could not be retrieved")
)

var (
	logFileExtension = ".log"
)

type AppendOnlyLog struct {
	// todo: add distributed name to the config (cluster instance)
	// todo: add lock function

	file *os.File
}

func NewLog(filePath string) (pool *AppendOnlyLog, err error) {
	e.InterruptIfEmpty(filePath)
	if strings.HasSuffix(filePath, logFileExtension) == false {
		filePath += logFileExtension
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	// todo: obtain pool lock for defined period in etcd.

	pool = &AppendOnlyLog{
		file: file,
	}

	err = pool.adjustToLastValidRecord()
	return
}

// todo: tests needed
func (pool *AppendOnlyLog) Write(record Record) (err error) {
	// Reserve space for record size.
	// (writing 4 bytes of data is atomic, so no partial write should occur).
	recordSizeBinary := []byte{0, 0, 0, 0}
	bytesWritten, err := pool.file.Write(recordSizeBinary)
	if bytesWritten != 4 {
		err = ErrPartialWriteOccurred
	}

	err = pool.file.Sync()
	if err != nil {
		return
	}

	// Write record after the space reserved.
	// In case if write would fail - empty value in record size space would indicate that
	// next record is broken.
	recordBinary, err := record.MarshalBinary()
	if err != nil {
		return
	}

	recordSize := len(recordBinary)
	bytesWritten, err = pool.file.Write(recordBinary)
	if err != nil {
		return
	}

	if bytesWritten != recordSize {
		err = ErrPartialWriteOccurred
		return
	}

	currentFilePos, err := pool.file.Seek(0, 1)
	if err != nil {
		return
	}

	// Writing record size on a previously reserved slot.
	// In case of success this would mark the record as valid and
	// available for reading on the next log parsing operation.
	_, err = pool.file.Seek(int64(recordSize+4)*(-1), 1)
	if err != nil {
		return
	}

	binary.BigEndian.PutUint32(recordSizeBinary, uint32(recordSize))
	bytesWritten, err = pool.file.Write(recordSizeBinary)
	if err != nil {
		return
	}

	if bytesWritten != 4 {
		err = ErrPartialWriteOccurred
		return
	}

	err = pool.file.Sync()
	if err != nil {
		return
	}

	_, err = pool.file.Seek(currentFilePos, 0)
	return
}

// todo: tests needed
func (pool *AppendOnlyLog) adjustToLastValidRecord() (err error) {
	recordSizeBinary := []byte{0, 0, 0, 0}

	for {
		// Reading first record.
		// In case if it fails - threat the log as empty and
		// seek to beginning (even if there is any info present)
		bytesRead, err := pool.file.Read(recordSizeBinary)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil // Not a problem.

			} else {
				return err
			}
		}

		if bytesRead == 0 {
			return nil // Not a problem. Just EOF
		}

		if bytesRead != len(recordSizeBinary) {
			// File seems to be corrupted (contains less info than even one record header).
			// In this case - no more important info could be present in it.
			// The logfile should be truncated up to the last record pos.
			lastValidRecordPos, err := pool.file.Seek(int64(-bytesRead), 1)
			if err != nil {
				return err
			}

			err = pool.file.Truncate(lastValidRecordPos)
			return err
		}

		recordSize := binary.BigEndian.Uint32(recordSizeBinary)
		if recordSize == 0 {
			// The write operation has finished partially.
			// It is ok. It means that during last record write there was a failure,
			// but the main structure of the log is ok.
			// In this case the last record should be simply skipped.

			lastValidRecordPos, err := pool.file.Seek(-4, 1)
			if err != nil {
				return err
			}

			err = pool.file.Truncate(lastValidRecordPos)
			return err
		}

		_, err = pool.file.Seek(int64(recordSize), 0)
		if err != nil {
			// File is corrupted: contains less data than reported record size.
			// File must not be changed or overwritten, because of a sensitive data it could store.
			// In this only report error and exit.
			err = ErrCorruptedLogFile
			return err
		}
	}
}
