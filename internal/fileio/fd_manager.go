package fileio

import (
	"os"
	"sync"
)

const (
	DefaultMaxFileNums = 256
)

const (
	TooManyFileOpenErrSuffix = "too many open files"
)

// FdManager hold a fd cache in memory, it lru based cache.
type FdManager struct {
	lock               sync.Mutex
	fdList             *doubleLinkedList
	size               int
	cleanThresholdNums int
	maxFdNums          int

	Cache map[string]*FdInfo
}

// NewFdm will return a fdManager object
func NewFdm(maxFdNums int, cleanThreshold float64) (fdm *FdManager) {
	_ = "STUB: not implemented"
	return nil
}

// FdInfo holds base fd info
type FdInfo struct {
	fd    *os.File
	path  string
	using uint
	next  *FdInfo
	prev  *FdInfo
}

func (fdInfo *FdInfo) Using() uint { _ = "STUB: not implemented"; return 0 }

// GetFd go through this method to get fd.
func (fdm *FdManager) GetFd(path string) (fd *os.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the numbers of fd in cache larger than the cleanThreshold in config, we will clean useless fd in cache

// if the numbers of fd in cache larger than the max numbers of fd in config, we will not add this fd to cache

// add this fd to cache

// determine if there are too many open files, we will first clean useless fd in cache and try open this file again

// if something wrong in cleanUselessFd, we will return "open too many files" err, because we want user not the main err is that

// try open this file again，if it still returns err, we will show this error to user

// add to cache if open this file successfully

// addToCache add fd to cache
func (fdm *FdManager) AddToCache(fd *os.File, cleanPath string) { _ = "STUB: not implemented"; return }

// reduceUsing when RWManager object close, it will go through this method let fdm know it return the fd to cache
func (fdm *FdManager) ReduceUsing(path string) { _ = "STUB: not implemented"; return }

// close means the cache.
func (fdm *FdManager) Close() error { _ = "STUB: not implemented"; return nil }

type doubleLinkedList struct {
	head *FdInfo
	tail *FdInfo
	size int
}

func initDoubleLinkedList() *doubleLinkedList { _ = "STUB: not implemented"; return nil }

func (list *doubleLinkedList) addNode(node *FdInfo) { _ = "STUB: not implemented"; return }

func (list *doubleLinkedList) removeNode(node *FdInfo) { _ = "STUB: not implemented"; return }

func (list *doubleLinkedList) moveNodeToFront(node *FdInfo) { _ = "STUB: not implemented"; return }

func (fdm *FdManager) cleanUselessFd() error { _ = "STUB: not implemented"; return nil }

func (fdm *FdManager) CloseByPath(path string) error { _ = "STUB: not implemented"; return nil }
