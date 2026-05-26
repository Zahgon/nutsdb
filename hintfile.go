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

package nutsdb

import (
	"bufio"
	"errors"
	"os"

	"github.com/nutsdb/nutsdb/internal/core"
)

const (
	// HintSuffix returns the hint file suffix
	HintSuffix = ".hint"
)

var (
	// ErrHintFileCorrupted is returned when hint file is corrupted
	ErrHintFileCorrupted = errors.New("hint file is corrupted")
	// ErrHintFileEntryInvalid is returned when hint entry is invalid
	ErrHintFileEntryInvalid = errors.New("hint file entry is invalid")
)

// getHintPath returns the hint file path for the given file ID and directory
func getHintPath(fid int64, dir string) string { _ = "STUB: not implemented"; return "" }

// HintEntry represents an entry in the hint file
type HintEntry struct {
	BucketId  uint64
	KeySize   uint32
	ValueSize uint32
	Timestamp uint64
	TTL       uint32
	Flag      uint16
	Status    uint16
	Ds        uint16
	DataPos   uint64
	FileID    int64
	Key       []byte
}

func newHintEntryFromEntry(entry *core.Entry, fileID int64, offset uint64) *HintEntry {
	_ = "STUB: not implemented"
	return nil
}

// Size returns the size of the hint entry
func (h *HintEntry) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Encode encodes the hint entry to bytes
func (h *HintEntry) Encode() []byte { _ = "STUB: not implemented"; return nil }

// Decode decodes the hint entry from bytes
func (h *HintEntry) Decode(buf []byte) error { _ = "STUB: not implemented"; return nil }

// HintFileReader is used to read hint files
type HintFileReader struct {
	file   *os.File
	reader *bufio.Reader
}

// Open opens a hint file for reading
func (r *HintFileReader) Open(path string) error { _ = "STUB: not implemented"; return nil }

// Read reads a hint entry from the file
func (r *HintFileReader) Read() (*HintEntry, error) { _ = "STUB: not implemented"; return nil, nil }

func decodeCompatInt64(buf []byte) (int64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func readCompatInt64(r *bufio.Reader, fieldsRead *int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close closes the hint file
func (r *HintFileReader) Close() error { _ = "STUB: not implemented"; return nil }

// HintFileWriter is used to write hint files
type HintFileWriter struct {
	file   *os.File
	writer *bufio.Writer
}

// Create creates a hint file for writing
func (w *HintFileWriter) Create(path string) error { _ = "STUB: not implemented"; return nil }

// Write writes a hint entry to the file
func (w *HintFileWriter) Write(entry *HintEntry) error { _ = "STUB: not implemented"; return nil }

// Sync flushes the buffer and syncs the file to disk
func (w *HintFileWriter) Sync() error { _ = "STUB: not implemented"; return nil }

// Close closes the hint file
func (w *HintFileWriter) Close() error { _ = "STUB: not implemented"; return nil }
