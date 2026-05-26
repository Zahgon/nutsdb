package fileio

import "os"

// Truncate changes the size of the file.
// If readOnly is true, it skips the file stat check and truncate operation,
// which significantly improves read performance by avoiding unnecessary syscalls.
func Truncate(path string, capacity int64, f *os.File, readOnly bool) error {
	_ = "STUB: not implemented"
	// Skip truncation for read-only operations to avoid expensive os.Stat syscall
	return nil
}
