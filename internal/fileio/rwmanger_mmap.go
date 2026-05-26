// Copyright 2019 The nutsdb Author. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileio

import (
	"errors"
	"os"
	"sync"

	"github.com/edsrzf/mmap-go"
	"github.com/nutsdb/nutsdb/internal/utils"
)

var (
	MmapBlockSize = int64(os.Getpagesize()) * 4

	mmapRWManagerInstancesLock = sync.Mutex{}
	mmapRWManagerInstances     = make(map[string]*MMapRWManager)
	mmapLRUCacheCapacity       = 1024
)

var (
	// ErrUnmappedMemory is returned when a function is called on unmapped memory
	ErrUnmappedMemory = errors.New("unmapped memory")

	// ErrIndexOutOfBound is returned when given offset out of mapped region
	ErrIndexOutOfBound = errors.New("offset out of mapped region")
)

func GetMMapRWManager(fd *os.File, path string, fdm *FdManager, segmentSize int64) *MMapRWManager {
	_ = "STUB: not implemented"
	return nil
}

// MMapRWManager represents the RWManager which using mmap.
// different with FileIORWManager, we can only allow one MMapRWManager
// for one datafile, so we need to do some concurrency safety control.
type MMapRWManager struct {
	Fd          *os.File
	Path        string
	Fdm         *FdManager
	SegmentSize int64

	ReadCache  *utils.LRUCache
	WriteCache *utils.LRUCache
}

// WriteAt copies data to mapped region from the b slice starting at
// given off and returns number of bytes copied to the mapped region.
func (mm *MMapRWManager) WriteAt(b []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadAt copies data to b slice from mapped region starting at
// given off and returns number of bytes copied to the b slice.
func (mm *MMapRWManager) ReadAt(b []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Sync synchronizes the mapping's contents to the file's contents on disk.
func (mm *MMapRWManager) Sync() (err error) {
	_ = "STUB: not implemented"

	// Release deletes the memory mapped region, flushes any remaining changes
	return nil
}

func (mm *MMapRWManager) Release() (err error) { _ = "STUB: not implemented"; return nil }

// Remove from global instances map to prevent reuse with stale FdManager

func (mm *MMapRWManager) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Close will remove the cache in the fdm of the specified path, and call the close method of the os of the file
func (mm *MMapRWManager) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (mm *MMapRWManager) alignedOffset(offset int64) int64 { _ = "STUB: not implemented"; return 0 }

// accessMMap access the MMap data. If for this block the mmap data is not mmapped, will add
// it into cache.
func (mm *MMapRWManager) accessMMap(cache *utils.LRUCache, offset int64, prot int) (data *mmapData, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mmapData is a struct to control the lifetime and access level of mmap.MMap
type mmapData struct {
	data   mmap.MMap
	offset int64
	mut    sync.RWMutex
}

func newMMapData(fd *os.File, offset int64, prot int) (md *mmapData, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (md *mmapData) Close() (err error) { _ = "STUB: not implemented"; return nil }
