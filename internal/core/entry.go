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

package core

import (
	"encoding/binary"
	"errors"
)

var (
	ErrPayLoadSizeMismatch   = errors.New("the payload size in Meta mismatch with the payload size needed")
	ErrHeaderSizeOutOfBounds = errors.New("the header size is out of bounds")

	// Error constants moved from root package
	ErrDataSizeExceed            = errors.New("data size too big")
	ErrKeyEmpty                  = errors.New("key cannot be empty")
	ErrInvalidKey                = errors.New("invalid key")
	ErrDataStructureNotSupported = errors.New("this data structure is not supported for now")
)

const (
	MaxEntryHeaderSize = 4 + binary.MaxVarintLen32*3 + binary.MaxVarintLen64*3 + binary.MaxVarintLen16*3
	MinEntryHeaderSize = 4 + 9

	// SeparatorForZSetKey represents separator for zSet key.
	SeparatorForZSetKey = "|"
)

type (
	// Entry represents the data item.
	Entry struct {
		Key   []byte
		Value []byte
		Meta  *MetaData
	}
)

// Size returns the size of the entry.
func (e *Entry) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Encode returns the slice after the entry be encoded.
//
//	the entry stored format:
//	|----------------------------------------------------------------------------------------------------------|
//	|  crc  | timestamp | ksz | valueSize | flag  | TTL  | status | ds   | txId |  bucketId |  key  | value    |
//	|----------------------------------------------------------------------------------------------------------|
//	| uint32| uint64  |uint32 |  uint32 | uint16  | uint32| uint16 | uint16 |uint64 | uint64 | []byte | []byte |
//	|----------------------------------------------------------------------------------------------------------|
func (e *Entry) Encode() []byte { _ = "STUB: not implemented"; return nil }

// setEntryHeaderBuf sets the entry header buff.
func (e *Entry) setEntryHeaderBuf(buf []byte) int { _ = "STUB: not implemented"; return 0 }

// IsZero checks if the entry is zero or not.
func (e *Entry) IsZero() bool { _ = "STUB: not implemented"; return false }

// GetCrc returns the crc at given buf slice.
func (e *Entry) GetCrc(buf []byte) uint32 { _ = "STUB: not implemented"; return 0 }

// ParsePayload means this function will parse a byte array to bucket, key, size of an entry
func (e *Entry) ParsePayload(data []byte) error { _ = "STUB: not implemented"; return nil }

// parse key

// parse value

// CheckPayloadSize checks the payload size
func (e *Entry) CheckPayloadSize(size int64) error { _ = "STUB: not implemented"; return nil }

// ParseMeta parse Meta object to entry
func (e *Entry) ParseMeta(buf []byte) (int64, error) {
	_ = "STUB: not implemented"
	// If the length of the header is less than MinEntryHeaderSize,
	// it means that the final remaining capacity of the file is not enough to write a record,
	// and an error needs to be returned.
	return 0, nil
}

// IsFilter to confirm if this entry is can be filtered
func (e *Entry) IsFilter() bool { _ = "STUB: not implemented"; return false }

// Valid check the entry fields valid or not
func (e *Entry) Valid() error { _ = "STUB: not implemented"; return nil }

// Note: MAX_SIZE will be re-exported from root package for backward compatibility
// For now, we'll use a reasonable default that works on both 32-bit and 64-bit systems
// 1GB, reasonable for most use cases

// NewEntry new Entry Object
func NewEntry() *Entry {
	_ = "STUB: not implemented"

	// WithKey set key to Entry
	return nil
}

func (e *Entry) WithKey(key []byte) *Entry { _ = "STUB: not implemented"; return nil }

// WithValue set value to Entry
func (e *Entry) WithValue(value []byte) *Entry { _ = "STUB: not implemented"; return nil }

// WithMeta set Meta to Entry
func (e *Entry) WithMeta(meta *MetaData) *Entry { _ = "STUB: not implemented"; return nil }

// GetTxIDBytes return the bytes of TxID
func (e *Entry) GetTxIDBytes() []byte { _ = "STUB: not implemented"; return nil }

func (e *Entry) IsBelongsToBTree() bool { _ = "STUB: not implemented"; return false }

// IsBelongsToBPlusTree is kept for backward compatibility with legacy naming.
// Internally nutsdb uses a B+ tree implementation for primary indexes, so both
// helpers map to the same metadata flag.
func (e *Entry) IsBelongsToBPlusTree() bool { _ = "STUB: not implemented"; return false }

func (e *Entry) IsBelongsToList() bool { _ = "STUB: not implemented"; return false }

func (e *Entry) IsBelongsToSet() bool { _ = "STUB: not implemented"; return false }

func (e *Entry) IsBelongsToSortSet() bool { _ = "STUB: not implemented"; return false }

type EntryWhenRecovery struct {
	Entry
	Fid int64
	Off int64
}

type DataInTx struct {
	Es       []*EntryWhenRecovery
	TxId     uint64
	StartOff int64
}

func (dt *DataInTx) IsSameTx(e *EntryWhenRecovery) bool { _ = "STUB: not implemented"; return false }

func (dt *DataInTx) AppendEntry(e *EntryWhenRecovery) { _ = "STUB: not implemented"; return }

func (dt *DataInTx) Reset() { _ = "STUB: not implemented"; return }

/**
 * decode the key of the entry
 * 1. in the case of list flag is DataLPushFlag or DataRPushFlag, the key is transformed from seq + user_key to user_key
 * so we need to decode the key to get the raw key
 * 2. in the case of sorted set flag is DataZAddFlag, the key is transformed from score + user_key to user_key
 * so we need to decode the key to get the raw key
 * 3. All other cases, the key is the raw key
 */
func (entry *Entry) GetRawKey() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
