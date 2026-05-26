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
	"github.com/nutsdb/nutsdb/internal/core"
)

const (
	getAllType    uint8 = 0
	getKeysType   uint8 = 1
	getValuesType uint8 = 2
)

func (tx *Tx) PutWithTimestamp(bucket string, key, value []byte, ttl uint32, timestamp uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Put sets the value for a key in the bucket.
// a wrapper of the function put.
func (tx *Tx) Put(bucket string, key, value []byte, ttl uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// PutIfNotExists set the value for a key in the bucket only if the key doesn't exist already.
func (tx *Tx) PutIfNotExists(bucket string, key, value []byte, ttl uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// the key-value is exists.

// PutIfExists set the value for a key in the bucket only if the key already exits.
func (tx *Tx) PutIfExists(bucket string, key, value []byte, ttl uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves the value for a key in the bucket.
// The returned value is only valid for the life of the transaction.
func (tx *Tx) Get(bucket string, key []byte) (value []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) get(bucket string, key []byte) (value []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Index layer automatically filters expired records

func (tx *Tx) ValueLen(bucket string, key []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) GetMaxKey(bucket string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *Tx) GetMinKey(bucket string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *Tx) getMaxOrMinKey(bucket string, isMax bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAll returns all keys and values in the given bucket.
func (tx *Tx) GetAll(bucket string) ([][]byte, [][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetKeys returns all keys in the given bucket.
func (tx *Tx) GetKeys(bucket string) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetValues returns all values in the given bucket.
func (tx *Tx) GetValues(bucket string) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) getAllOrKeysOrValues(bucket string, typ uint8) ([][]byte, [][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (tx *Tx) GetSet(bucket string, key, value []byte) (oldValue []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Has returns true if the record exists. It checks without retrieving the
// value from disk making lookups significantly faster while keeping memory
// usage down as well. It does require the `HintKeyAndRAMIdxMode` option to be
// enabled to function as described.
func (tx *Tx) Has(bucket string, key []byte) (exists bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RangeScanEntries query a range at given bucket, start and end slice. It will
// return keys and/or values based on the includeKeys and includeValues flags.
func (tx *Tx) RangeScanEntries(bucket string, start, end []byte, includeKeys, includeValues bool) (keys, values [][]byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If there is no item in pending and persist db,
// return error itself.

// Check for empty results before setting to nil

// RangeScan query a range at given bucket, start and end slice.
func (tx *Tx) RangeScan(bucket string, start, end []byte) (values [][]byte, err error) {
	_ = "STUB: not implemented"
	// RangeScan is kept as an API call to not break upstream projects that
	// rely on it.
	return nil, nil
}

// PrefixScanEntries iterates over a key prefix at given bucket, prefix and
// limitNum.  If reg is set a regular expression will be used to filter the
// found entries. LimitNum will limit the number of entries return. It will
// return keys and/or values based on the includeKeys and includeValues flags.
func (tx *Tx) PrefixScanEntries(bucket string, prefix []byte, reg string, offsetNum int, limitNum int, includeKeys, includeValues bool) (keys, values [][]byte, err error) {
	_ = "STUB: not implemented"
	// This function is a bit awkward but that is to maintain backwards
	// compatibility while enabling the caller to pick and choose which
	// variation to call.
	return nil, nil, nil
}

// Return expected error types based on Scan/SearchScan.

// PrefixScan iterates over a key prefix at given bucket, prefix and limitNum.
// LimitNum will limit the number of entries return.
func (tx *Tx) PrefixScan(bucket string, prefix []byte, offsetNum int, limitNum int) (values [][]byte, err error) {
	_ = "STUB: not implemented"
	// PrefixScan is kept as an API call to not break upstream projects
	// that rely on it.
	return nil, nil
}

// PrefixSearchScan iterates over a key prefix at given bucket, prefix, match regular expression and limitNum.
// LimitNum will limit the number of entries return.
func (tx *Tx) PrefixSearchScan(bucket string, prefix []byte, reg string, offsetNum int, limitNum int) (values [][]byte, err error) {
	_ = "STUB: not implemented"
	// PrefixSearchScan is kept as an API call to not break upstream projects
	// that rely on it.
	return nil, nil
}

// Delete removes a key from the bucket at given bucket and key.
func (tx *Tx) Delete(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }

//Find the key in the transaction-local map (entriesInBTree)

// getHintIdxDataItemsWrapper returns keys and values when prefix scanning or range scanning.
// Note: TTL filtering is handled at the index layer, so records passed here are already valid (non-expired).
func (tx *Tx) getHintIdxDataItemsWrapper(records []*core.Record, limitNum int, _ core.BucketId, needKeys bool, needValues bool) (keys [][]byte, values [][]byte, err error) {
	_ = "STUB: not implemented"
	// Pre-allocate capacity to reduce slice re-growth
	return nil, nil, nil
}

func (tx *Tx) tryGet(bucket string, key []byte, solveRecord func(record *core.Record, entry *core.Entry, found bool, bucketId core.BucketId) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) update(bucket string, key []byte, getNewValue func([]byte) ([]byte, error), getNewTTL func(uint32) (uint32, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) updateOrPut(bucket string, key, value []byte, getUpdatedValue func([]byte) ([]byte, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func bigIntIncr(a string, b string) string { _ = "STUB: not implemented"; return "" }

func (tx *Tx) integerIncr(bucket string, key []byte, increment int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) Incr(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) Decr(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) IncrBy(bucket string, key []byte, increment int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) DecrBy(bucket string, key []byte, decrement int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) GetBit(bucket string, key []byte, offset int) (byte, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) SetBit(bucket string, key []byte, offset int, bit byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTTL returns remaining TTL of a value by key.
// It returns
// (-1, nil) If TTL is Persistent
// (0, ErrBucketNotFound|ErrKeyNotFound) If expired or not found
// (TTL, nil) If the record exists with a TTL
// Note: The returned remaining TTL will be in seconds. For example,
// remainingTTL is 500ms, It'll return 0.
func (tx *Tx) GetTTL(bucket string, key []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check pending writes first

// Persist updates record's TTL as Persistent if the record exists.
func (tx *Tx) Persist(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) MSet(bucket string, ttl uint32, args ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) MGet(bucket string, keys ...[]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) Append(bucket string, key, appendage []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) GetRange(bucket string, key []byte, start, end int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadValue in order to load value during pending and
// stored items, we need this function to load value and
// TTL.
func (tx *Tx) loadValue(
	rec *core.Record,
	pendingEntry *core.Entry,
) (value []byte, ttl uint32, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
