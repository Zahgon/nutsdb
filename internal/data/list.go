package data

import (
	"errors"

	"github.com/nutsdb/nutsdb/internal/core"
)

var (
	// ErrListNotFound is returned when the list not found.
	ErrListNotFound = errors.New("the list not found")

	// ErrCount is returned when count is error.
	ErrCount = errors.New("err count")

	// ErrEmptyList is returned when the list is empty.
	ErrEmptyList = errors.New("the list is empty")

	// ErrStartOrEnd is returned when start > end
	ErrStartOrEnd = errors.New("start or end error")
)

type ListImplementationType int

const (
	// ListImplDoublyLinkedList uses doubly linked list implementation (default).
	// Advantages: O(1) head/tail operations, lower memory overhead
	// Best for: High-frequency LPush/RPush/LPop/RPop operations
	ListImplDoublyLinkedList = iota

	// ListImplBTree uses BTree implementation.
	// Advantages: O(log n + k) range queries, efficient random access
	// Best for: Frequent range queries or indexed access patterns
	ListImplBTree
)

// HeadTailSeq list head and tail seq num
type HeadTailSeq struct {
	Head uint64
	Tail uint64
}

func (seq *HeadTailSeq) GenerateSeq(isLeft bool) uint64 { _ = "STUB: not implemented"; return 0 }

// ListStructure defines the interface for List storage implementations.
// It supports multiple implementations: BTree, DoublyLinkedList, SkipList, etc.
// This interface enables users to choose the most suitable implementation based on their use case:
// - DoublyLinkedList: O(1) head/tail operations, optimal for LPush/RPush/LPop/RPop
// - BTree: O(log n) operations, better for range queries and random access
type ListStructure interface {
	// InsertRecord inserts a record with the given key (sequence number in big-endian format).
	// Returns true if an existing record was replaced, false if a new record was inserted.
	InsertRecord(key []byte, record *core.Record) bool

	// Delete removes the record with the given key.
	// Returns true if the record was found and deleted, false otherwise.
	Delete(key []byte) bool

	// Find retrieves the record with the given key.
	// Returns the record and true if found, nil and false otherwise.
	Find(key []byte) (*core.Record, bool)

	// Min returns the item with the smallest key (head of the list).
	// Returns the item and true if the list is not empty, nil and false otherwise.
	Min() (*core.Item[core.Record], bool)

	// Max returns the item with the largest key (tail of the list).
	// Returns the item and true if the list is not empty, nil and false otherwise.
	Max() (*core.Item[core.Record], bool)

	// All returns all records in ascending key order.
	All() []*core.Record

	// AllItems returns all items (key + record pairs) in ascending key order.
	AllItems() []*core.Item[core.Record]

	// Count returns the number of elements in the list.
	Count() int

	// Range returns records with keys in the range [start, end] (inclusive).
	Range(start, end []byte) []*core.Record

	// PrefixScan scans records with keys matching the given prefix.
	// offset: number of matching records to skip
	// limitNum: maximum number of records to return
	PrefixScan(prefix []byte, offset, limitNum int) []*core.Record

	// PrefixSearchScan scans records with keys matching the given prefix and regex pattern.
	// The regex is applied to the portion of the key after removing the prefix.
	// offset: number of matching records to skip
	// limitNum: maximum number of records to return
	PrefixSearchScan(prefix []byte, reg string, offset, limitNum int) []*core.Record

	// PopMin removes and returns the item with the smallest key.
	// Returns the item and true if the list is not empty, nil and false otherwise.
	PopMin() (*core.Item[core.Record], bool)

	// PopMax removes and returns the item with the largest key.
	// Returns the item and true if the list is not empty, nil and false otherwise.
	PopMax() (*core.Item[core.Record], bool)
}

// Compile-time interface implementation checks
var (
	_ ListStructure = (*BTree)(nil)
	_ ListStructure = (*DoublyLinkedList)(nil)
)

// BTree represents the btree.

// List represents the list.
type List struct {
	Items     map[string]ListStructure
	TTL       map[string]uint32
	TimeStamp map[string]uint64
	Seq       map[string]*HeadTailSeq
	ListImpl  ListImplementationType
}

func NewList(listImpl ListImplementationType) *List { _ = "STUB: not implemented"; return nil }

// CreateListStructure creates a new list storage structure based on configuration.
func (l *List) CreateListStructure() ListStructure {
	_ = "STUB: not implemented"
	return *new(ListStructure)
}

// Default to DoublyLinkedList for safety

func (l *List) LPush(key string, r *core.Record) error { _ = "STUB: not implemented"; return nil }

func (l *List) RPush(key string, r *core.Record) error { _ = "STUB: not implemented"; return nil }

func (l *List) Push(key string, r *core.Record, isLeft bool) error {
	_ = "STUB: not implemented"
	// key is seq + user_key
	return nil
}

// Initialize seq if not exists

// Update seq boundaries to track the next insertion positions
// This is important for recovery scenarios where we rebuild the index
// Head and Tail should always represent the next available positions for insertion

// LPush: Head should be the next available position on the left
// If current seq is the actual head, set Head to current seq - 1

// RPush: Tail should be the next available position on the right
// If current seq is at or beyond current tail, update Tail accordingly

func (l *List) LPop(key string) (*core.Record, error) { _ = "STUB: not implemented"; return nil, nil }

// Use PopMin for efficient O(1) head removal

// After LPop, Head should point to the next element's position
// Note: We don't update Head here because it represents "next push position"
// The popped element's sequence is already consumed

// RPop removes and returns the last element of the list stored at key.
func (l *List) RPop(key string) (*core.Record, error) { _ = "STUB: not implemented"; return nil, nil }

// Use PopMax for efficient O(1) tail removal

// After RPop, Tail should point to the next element's position
// Note: We don't update Tail here because it represents "next push position"
// The popped element's sequence is already consumed

func (l *List) LPeek(key string) (*core.Item[core.Record], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *List) RPeek(key string) (*core.Item[core.Record], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *List) peek(key string, isLeft bool) (*core.Item[core.Record], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LRange returns the specified elements of the list stored at key [start,end]
func (l *List) LRange(key string, start, end int) ([]*core.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRemoveIndexes returns a slice of indices to be removed from the list based on the count
func (l *List) GetRemoveIndexes(key string, count int, cmp func(r *core.Record) (bool, error)) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LRem removes the first count occurrences of elements equal to value from the list stored at key.
// The count argument influences the operation in the following ways:
// count > 0: Remove elements equal to value moving from head to tail.
// count < 0: Remove elements equal to value moving from tail to head.
// count = 0: Remove all elements equal to value.
func (l *List) LRem(key string, count int, cmp func(r *core.Record) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// LTrim trim an existing list so that it will contain only the specified range of elements specified.
func (l *List) LTrim(key string, start, end int) error { _ = "STUB: not implemented"; return nil }

// LRemByIndex remove the list element at specified index
func (l *List) LRemByIndex(key string, indexes []int) error { _ = "STUB: not implemented"; return nil }

func (l *List) GetValidIndexes(key string, indexes []int) map[int]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (l *List) IsExpire(key string) bool { _ = "STUB: not implemented"; return false }

func (l *List) Size(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *List) IsEmpty(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (l *List) GetListTTL(key string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *List) ExpireList(key []byte, ttl uint32) { _ = "STUB: not implemented"; return }

func (l *List) GeneratePushKey(key []byte, isLeft bool) []byte {
	_ = "STUB: not implemented"
	// Retrieve or initialize the HeadTailSeq for the list
	return nil
}

// If no seq entry exists, infer boundaries from existing items first

func checkBounds(start, end int, size int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
