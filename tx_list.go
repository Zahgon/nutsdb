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
	"github.com/nutsdb/nutsdb/internal/data"
	"github.com/pkg/errors"
)

// SeparatorForListKey represents separator for listKey
const SeparatorForListKey = "|"

var (
	// ErrListNotFound is returned when the list not found.
	ErrListNotFound = data.ErrListNotFound

	// ErrCount is returned when count is error.
	ErrCount = data.ErrCount

	// ErrEmptyList is returned when the list is empty.
	ErrEmptyList = data.ErrEmptyList

	// ErrStartOrEnd is returned when start > end
	ErrStartOrEnd = data.ErrStartOrEnd

	// ErrSeparatorForListKey returns when list key contains the SeparatorForListKey.
	ErrSeparatorForListKey = errors.Errorf("contain separator (%s) for List key", SeparatorForListKey)
)

// RPop removes and returns the last element of the list stored in the bucket at given bucket and key.
func (tx *Tx) RPop(bucket string, key []byte) (item []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RPeek returns the last element of the list stored in the bucket at given bucket and key.
func (tx *Tx) RPeek(bucket string, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// push sets values for list stored in the bucket at given bucket, key, flag and values.
func (tx *Tx) push(bucket string, key []byte, flag uint16, values ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// getListWithDefault
// this function will get list, if list not exists, will create
// a new one.
func (tx *Tx) getListWithDefault(bucket string) (*data.List, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure the list index exists before performing list operations

// RPush inserts the values at the tail of the list stored in the bucket at given bucket,key and values.
func (tx *Tx) RPush(bucket string, key []byte, values ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// LPush inserts the values at the head of the list stored in the bucket at given bucket,key and values.
func (tx *Tx) LPush(bucket string, key []byte, values ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) isKeyValid(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) LPushRaw(bucket string, key []byte, values ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) RPushRaw(bucket string, key []byte, values ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// LPop removes and returns the first element of the list stored in the bucket at given bucket and key.
func (tx *Tx) LPop(bucket string, key []byte) (item []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LPeek returns the first element of the list stored in the bucket at given bucket and key.
func (tx *Tx) LPeek(bucket string, key []byte) (item []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LSize returns the size of key in the bucket in the bucket at given bucket and key.
func (tx *Tx) LSize(bucket string, key []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// LRange returns the specified elements of the list stored in the bucket at given bucket,key, start and end.
// The offsets start and stop are zero-based indexes 0 being the first element of the list (the head of the list),
// 1 being the next element and so on.
// Start and end can also be negative numbers indicating offsets from the end of the list,
// where -1 is the last element of the list, -2 the penultimate element and so on.
func (tx *Tx) LRange(bucket string, key []byte, start, end int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LRem removes the first count occurrences of elements equal to value from the list stored in the bucket at given bucket,key,count.
// The count argument influences the operation in the following ways:
// count > 0: Remove elements equal to value moving from head to tail.
// count < 0: Remove elements equal to value moving from tail to head.
// count = 0: Remove all elements equal to value.
func (tx *Tx) LRem(bucket string, key []byte, count int, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// LTrim trims an existing list so that it will contain only the specified range of elements specified.
// the offsets start and stop are zero-based indexes 0 being the first element of the list (the head of the list),
// 1 being the next element and so on.
// start and end can also be negative numbers indicating offsets from the end of the list,
// where -1 is the last element of the list, -2 the penultimate element and so on.
func (tx *Tx) LTrim(bucket string, key []byte, start, end int) error {
	_ = "STUB: not implemented"
	return nil
}

// LRemByIndex remove the list element at specified index
func (tx *Tx) LRemByIndex(bucket string, key []byte, indexes ...int) error {
	_ = "STUB: not implemented"
	return nil
}

// LKeys find all keys matching a given pattern
func (tx *Tx) LKeys(bucket, pattern string, f func(key string) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) ExpireList(bucket string, key []byte, ttl uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) CheckExpire(bucket string, key []byte) bool { _ = "STUB: not implemented"; return false }

func (tx *Tx) GetListTTL(bucket string, key []byte) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
