// Copyright 2023 The nutsdb Author. All rights reserved.
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

package data

import (
	"errors"

	"github.com/nutsdb/nutsdb/internal/core"
)

var (
	// ErrSetNotExist is returned when the key does not exist.
	ErrSetNotExist = errors.New("set not exist")

	// ErrSetMemberNotExist is returned when the member of set does not exist
	ErrSetMemberNotExist = errors.New("set member not exist")

	// ErrMemberEmpty is returned when the item received is nil
	ErrMemberEmpty = errors.New("item empty")
)

type Set struct {
	M map[string]map[uint32]*core.Record
}

func NewSet() *Set { _ = "STUB: not implemented"; return nil }

// SAdd adds the specified members to the set stored at key.
func (s *Set) SAdd(key string, values [][]byte, records []*core.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// SRem removes the specified members from the set stored at key.
func (s *Set) SRem(key string, values ...[]byte) error { _ = "STUB: not implemented"; return nil }

// SHasKey returns whether it has the set at given key.
func (s *Set) SHasKey(key string) bool { _ = "STUB: not implemented"; return false }

// SPop removes and returns one or more random elements from the set value store at key.
func (s *Set) SPop(key string) *core.Record { _ = "STUB: not implemented"; return nil }

// SCard Returns the set cardinality (number of elements) of the set stored at key.
func (s *Set) SCard(key string) int { _ = "STUB: not implemented"; return 0 }

// SDiff Returns the members of the set resulting from the difference between the first set and all the successive sets.
func (s *Set) SDiff(key1, key2 string) ([]*core.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SInter Returns the members of the set resulting from the intersection of all the given sets.
func (s *Set) SInter(key1, key2 string) ([]*core.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SIsMember Returns if member is a member of the set stored at key.
func (s *Set) SIsMember(key string, value []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SAreMembers Returns if members are members of the set stored at key.
// For multiple items it returns true only if all the items exist.
func (s *Set) SAreMembers(key string, values ...[]byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SMembers returns all the members of the set value stored at key.
func (s *Set) SMembers(key string) ([]*core.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SMove moves member from the set at source to the set at destination.
func (s *Set) SMove(key1, key2 string, value []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SUnion returns the members of the set resulting from the union of all the given sets.
func (s *Set) SUnion(key1, key2 string) ([]*core.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
