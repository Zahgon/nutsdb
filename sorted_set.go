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

package nutsdb

import (
	"errors"

	"github.com/nutsdb/nutsdb/internal/core"
)

var (
	ErrSortedSetNotFound = errors.New("the sortedSet does not exist")

	ErrSortedSetMemberNotExist = errors.New("the member of sortedSet does not exist")

	ErrSortedSetIsEmpty = errors.New("the sortedSet if empty")
)

const (
	// SkipListMaxLevel represents the skipList max level number.
	SkipListMaxLevel = 32

	// SkipListP represents the p parameter of the skipList.
	SkipListP = 0.25
)

type SortedSet struct {
	db *DB
	M  map[string]*SkipList
}

func NewSortedSet(db *DB) *SortedSet { _ = "STUB: not implemented"; return nil }

func (z *SortedSet) ZAdd(key string, score SCORE, value []byte, record *core.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (z *SortedSet) ZMembers(key string) (map[*core.Record]SCORE, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *SortedSet) ZCard(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (z *SortedSet) ZCount(key string, start SCORE, end SCORE, opts *GetByScoreRangeOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (z *SortedSet) ZPeekMax(key string) (*core.Record, SCORE, error) {
	_ = "STUB: not implemented"
	return nil, *new(SCORE), nil
}

func (z *SortedSet) ZPopMax(key string) (*core.Record, SCORE, error) {
	_ = "STUB: not implemented"
	return nil, *new(SCORE), nil
}

func (z *SortedSet) ZPeekMin(key string) (*core.Record, SCORE, error) {
	_ = "STUB: not implemented"
	return nil, *new(SCORE), nil
}

func (z *SortedSet) ZPopMin(key string) (*core.Record, SCORE, error) {
	_ = "STUB: not implemented"
	return nil, *new(SCORE), nil
}

func (z *SortedSet) ZRangeByScore(key string, start SCORE, end SCORE, opts *GetByScoreRangeOptions) ([]*core.Record, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (z *SortedSet) ZRangeByRank(key string, start int, end int) ([]*core.Record, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (z *SortedSet) ZRem(key string, value []byte) (*core.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *SortedSet) ZRemRangeByRank(key string, start int, end int) error {
	_ = "STUB: not implemented"
	return nil
}

func (z *SortedSet) getZRemRangeByRankNodes(key string, start int, end int) ([]*SkipListNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *SortedSet) ZRank(key string, value []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (z *SortedSet) ZRevRank(key string, value []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (z *SortedSet) ZScore(key string, value []byte) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (z *SortedSet) ZExist(key string, value []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SCORE represents the score type.
type SCORE float64

// SkipListLevel records forward and span.
type SkipListLevel struct {
	forward *SkipListNode
	span    int64
}

// The SkipList represents the sorted set.
type SkipList struct {
	db     *DB
	header *SkipListNode
	tail   *SkipListNode
	length int64
	level  int
	dict   map[uint32]*SkipListNode
}

// SkipListNode represents a node in the SkipList.
type SkipListNode struct {
	hash     uint32       // unique key of this node
	record   *core.Record // associated data
	score    SCORE        // score to determine the order of this node in the set
	backward *SkipListNode
	level    []SkipListLevel
}

// Hash returns the key of the node.
func (sln *SkipListNode) Hash() uint32 {
	_ = "STUB: not implemented"

	// Score returns the score of the node.
	return 0
}

func (sln *SkipListNode) Score() SCORE {
	_ = "STUB: not implemented"

	// createNode returns a newly initialized SkipListNode Object that implements the SkipListNode.
	return *new(SCORE)
}

func createNode(level int, score SCORE, hash uint32, record *core.Record) *SkipListNode {
	_ = "STUB: not implemented"
	return nil
}

// randomLevel returns a random level for the new skiplist node we are going to create.
// The return value of this function is between 1 and SkipListMaxLevel
// (both inclusive), with a powerlaw-alike distribution where higher
// levels are lesl likely to be returned.
func randomLevel() int { _ = "STUB: not implemented"; return 0 }

func newSkipList(db *DB) *SkipList { _ = "STUB: not implemented"; return nil }

func (sl *SkipList) cmp(r1 *core.Record, r2 *core.Record) int { _ = "STUB: not implemented"; return 0 }

func (sl *SkipList) insertNode(score SCORE, hash uint32, record *core.Record) *SkipListNode {
	_ = "STUB: not implemented"
	return nil
}

// store rank that is crosled to reach the insert position

// score is the same but the key is different

/* we assume the key is not already inside, since we allow duplicated
 * scores, and the re-insertion of score and redis object should never
 * happen since the caller of Insert() should test in the hash table
 * if the element is already inside or not. */

// add a new level

/* update span covered by update[i] as x is inserted here */

// increment span for untouched levels

// deleteNode represents internal function used by delete, DeleteByScore and DeleteByRank.
func (sl *SkipList) deleteNode(x *SkipListNode, update [SkipListMaxLevel]*SkipListNode) {
	_ = "STUB: not implemented"
	return
}

// delete removes an element with matching score/key from the skiplist.
func (sl *SkipList) delete(score SCORE, hash uint32) bool { _ = "STUB: not implemented"; return false }

/* We may have multiple elements with the same score, what we need
 * is to find the element with both the right score and object. */

// free x

/* not found */

// Size returns the number of elements in the SkipList.
func (sl *SkipList) Size() int { _ = "STUB: not implemented"; return 0 }

// PeekMin returns the element with minimum score, nil if the set is empty.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) PeekMin() *SkipListNode { _ = "STUB: not implemented"; return nil }

// PopMin returns and remove the element with minimal score, nil if the set is empty.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) PopMin() *SkipListNode { _ = "STUB: not implemented"; return nil }

// PeekMax returns the element with maximum score, nil if the set is empty.
//
// Time Complexity : O(1).
func (sl *SkipList) PeekMax() *SkipListNode {
	_ = "STUB: not implemented"

	// PopMax returns and remove the element with maximum score, nil if the set is empty.
	//
	// Time complexity of this method is : O(log(N)).
	return nil
}

func (sl *SkipList) PopMax() *SkipListNode { _ = "STUB: not implemented"; return nil }

// Put puts an element into the sorted set with specific key / value / score.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) Put(score SCORE, value []byte, record *core.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// score does not change, only update value
// score changes, delete and re-insert

// Remove removes element specified at given key.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) Remove(hash uint32) *SkipListNode { _ = "STUB: not implemented"; return nil }

// GetByScoreRangeOptions represents the options of the GetByScoreRange function.
type GetByScoreRangeOptions struct {
	Limit        int  // limit the max nodes to return
	ExcludeStart bool // exclude start value, so it search in interval (start, end] or (start, end)
	ExcludeEnd   bool // exclude end value, so it search in interval [start, end) or (start, end)
}

// GetByScoreRange returns the nodes whose score within the specific range.
// If options is nil, it searches in interval [start, end] without any limit by default.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) GetByScoreRange(start SCORE, end SCORE, options *GetByScoreRangeOptions) []*SkipListNode {
	_ = "STUB: not implemented"
	return nil
}

// determine if out of range

// search from end to start

// search from start to end

func (sl *SkipList) searchForward(nodes []*SkipListNode, excludeStart, excludeEnd bool, start, end SCORE, limit int) []*SkipListNode {
	_ = "STUB: not implemented"
	// search from start to end
	return nil
}

/* Current node is the last with score < or <= start. */

func (sl *SkipList) searchReverse(nodes []*SkipListNode, excludeStart, excludeEnd bool, start, end SCORE, limit int) []*SkipListNode {
	_ = "STUB: not implemented"
	return nil
}

// GetByRankRange returns nodes within specific rank range [start, end].
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node
// If start is greater than end, the returned array is in reserved order
// If remove is true, the returned nodes are removed.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) GetByRankRange(start, end int, remove bool) []*SkipListNode {
	_ = "STUB: not implemented"
	return nil
}

// swap start and end

func (sl *SkipList) sanitizeIndexes(start, end int) (newStart, newEnd int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// GetByRank returns the node at given rank.
// Note that the rank is 1-based integer. Rank 1 means the first node; Rank -1 means the last node.
// If remove is true, the returned nodes are removed
// If node is not found at specific rank, nil is returned.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) GetByRank(rank int, remove bool) *SkipListNode {
	_ = "STUB: not implemented"
	return nil
}

// GetByValue returns the node at given key.
// If node is not found, nil is returned
//
// Time complexity : O(1).
func (sl *SkipList) GetByValue(value []byte) *SkipListNode { _ = "STUB: not implemented"; return nil }

// FindRank Returns the rank of member in the sorted set stored at key, with the scores ordered from low to high.
// Note that the rank is 1-based integer. Rank 1 means the first node
// If the node is not found, 0 is returned. Otherwise rank(> 0) is returned.
//
// Time complexity of this method is : O(log(N)).
func (sl *SkipList) FindRank(hash uint32) int { _ = "STUB: not implemented"; return 0 }

// FindRevRank Returns the rank of member in the sorted set stored at key, with the scores ordered from high to low.
func (sl *SkipList) FindRevRank(hash uint32) int { _ = "STUB: not implemented"; return 0 }
