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
)

var (
	// ErrSetNotExist is returned when the key does not exist.
	ErrSetNotExist = data.ErrSetNotExist

	// ErrSetMemberNotExist is returned when the member of set does not exist
	ErrSetMemberNotExist = data.ErrSetMemberNotExist

	// ErrMemberEmpty is returned when the item received is nil
	ErrMemberEmpty = data.ErrMemberEmpty
)

func (tx *Tx) sPut(bucket string, key []byte, dataFlag uint16, values ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SAdd adds the specified members to the set stored int the bucket at given bucket,key and items.
func (tx *Tx) SAdd(bucket string, key []byte, items ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SRem removes the specified members from the set stored int the bucket at given bucket,key and items.
func (tx *Tx) SRem(bucket string, key []byte, items ...[]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SAreMembers returns if the specified members are the member of the set int the bucket at given bucket,key and items.
func (tx *Tx) SAreMembers(bucket string, key []byte, items ...[]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SIsMember returns if member is a member of the set stored int the bucket at given bucket,key and item.
func (tx *Tx) SIsMember(bucket string, key, item []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SMembers returns all the members of the set value stored int the bucket at given bucket and key.
func (tx *Tx) SMembers(bucket string, key []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SHasKey returns if the set in the bucket at given bucket and key.
func (tx *Tx) SHasKey(bucket string, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SPop removes and returns one or more random elements from the set value store in the bucket at given bucket and key.
func (tx *Tx) SPop(bucket string, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SCard returns the set cardinality (number of elements) of the set stored in the bucket at given bucket and key.
func (tx *Tx) SCard(bucket string, key []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SDiffByOneBucket returns the members of the set resulting from the difference
// between the first set and all the successive sets in one bucket.
func (tx *Tx) SDiffByOneBucket(bucket string, key1, key2 []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SDiffByTwoBuckets returns the members of the set resulting from the difference
// between the first set and all the successive sets in two buckets.
func (tx *Tx) SDiffByTwoBuckets(bucket1 string, key1 []byte, bucket2 string, key2 []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SMoveByOneBucket moves member from the set at source to the set at destination in one bucket.
func (tx *Tx) SMoveByOneBucket(bucket string, key1, key2, item []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SMoveByTwoBuckets moves member from the set at source to the set at destination in two buckets.
func (tx *Tx) SMoveByTwoBuckets(bucket1 string, key1 []byte, bucket2 string, key2, item []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SUnionByOneBucket the members of the set resulting from the union of all the given sets in one bucket.
func (tx *Tx) SUnionByOneBucket(bucket string, key1, key2 []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SUnionByTwoBuckets the members of the set resulting from the union of all the given sets in two buckets.
func (tx *Tx) SUnionByTwoBuckets(bucket1 string, key1 []byte, bucket2 string, key2 []byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SKeys find all keys matching a given pattern
func (tx *Tx) SKeys(bucket, pattern string, f func(key string) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrBucketAndKey returns when bucket or key not found.
func ErrBucketAndKey(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }

// ErrNotFoundKeyInBucket returns when key not in the bucket.
func ErrNotFoundKeyInBucket(bucket string, key []byte) error { _ = "STUB: not implemented"; return nil }
