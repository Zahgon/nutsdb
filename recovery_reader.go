package nutsdb

import (
	"bufio"
	"os"

	"github.com/nutsdb/nutsdb/internal/core"
)

// fileRecovery use bufio.Reader to read entry
type fileRecovery struct {
	fd     *os.File
	reader *bufio.Reader
	size   int64
	closed bool
}

func newFileRecovery(path string, bufSize int) (fr *fileRecovery, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readEntry will read an Entry from disk.
func (fr *fileRecovery) readEntry(off int64) (e *core.Entry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Since core.MaxEntryHeaderSize may be larger than the actual Header, it needs to be calculated

func (fr *fileRecovery) readBucket() (b *core.Bucket, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// calBufferSize calculates the buffer size of bufio.Reader
// if the size < 4 * KB, use 4 * KB as the size of buffer in bufio.Reader
// if the size > 4 * KB, use the nearly blockSize buffer as the size of buffer in bufio.Reader
func calBufferSize(size int) int { _ = "STUB: not implemented"; return 0 }

func (fr *fileRecovery) release() error { _ = "STUB: not implemented"; return nil }
