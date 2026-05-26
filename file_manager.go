package nutsdb

import "github.com/nutsdb/nutsdb/internal/fileio"

// RWMode represents the read and write mode.
type RWMode int

const (
	// FileIO represents the read and write mode using standard I/O.
	FileIO RWMode = iota

	// MMap represents the read and write mode using mmap.
	MMap
)

// FileManager holds the fd cache and file-related operations go through the manager to obtain the file processing object
type FileManager struct {
	rwMode      RWMode
	fdm         *fileio.FdManager
	segmentSize int64
}

// NewFileManager will create a NewFileManager object
func NewFileManager(rwMode RWMode, maxFdNums int, cleanThreshold float64, segmentSize int64) (fm *FileManager) {
	_ = "STUB: not implemented"
	return nil
}

// GetDataFile will return a DataFile Object
func (fm *FileManager) GetDataFile(path string, capacity int64) (datafile *DataFile, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDataFileReadOnly will return a DataFile Object for read-only operations
// This method skips file truncation to improve read performance
func (fm *FileManager) GetDataFileReadOnly(path string, capacity int64) (datafile *DataFile, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getDataFileWithMode will return a DataFile Object with specified read-only mode
func (fm *FileManager) getDataFileWithMode(path string, capacity int64, readOnly bool) (datafile *DataFile, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fm *FileManager) GetDataFileByID(dir string, fileID int64, capacity int64) (*DataFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFileRWManager will return a FileIORWManager Object
func (fm *FileManager) GetFileRWManager(path string, capacity int64, segmentSize int64, readOnly bool) (*fileio.FileIORWManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMMapRWManager will return a MMapRWManager Object
func (fm *FileManager) GetMMapRWManager(path string, capacity int64, segmentSize int64, readOnly bool) (*fileio.MMapRWManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close will Close fdm resource
func (fm *FileManager) Close() error { _ = "STUB: not implemented"; return nil }
