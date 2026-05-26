package utils

import (
	"container/list"
	"sync"
)

// LRUCache is a least recently used (LRU) cache.
type LRUCache struct {
	m   map[any]*list.Element
	l   *list.List
	cap int
	mu  *sync.RWMutex
}

// New creates a new LRUCache with the specified capacity.
func NewLruCache(cap int) *LRUCache { _ = "STUB: not implemented"; return nil }

// Add adds a new entry to the cache.
func (c *LRUCache) Add(key any, value any) { _ = "STUB: not implemented"; return }

// Get returns the entry associated with the given key, or nil if the key is not in the cache.
func (c *LRUCache) Get(key any) any { _ = "STUB: not implemented"; return *new(any) }

// Remove removes the entry associated with the given key from the cache.
func (c *LRUCache) Remove(key any) { _ = "STUB: not implemented"; return }

// Len returns the number of entries in the cache.
func (c *LRUCache) Len() int { _ = "STUB: not implemented"; return 0 }

// Clear clears the cache.
func (c *LRUCache) Clear() { _ = "STUB: not implemented"; return }

// removeOldest removes the oldest entry from the cache.
func (c *LRUCache) removeOldest() { _ = "STUB: not implemented"; return }

// LruEntry is a struct that represents an entry in the LRU cache.
type LruEntry struct {
	Key   any
	Value any
}
