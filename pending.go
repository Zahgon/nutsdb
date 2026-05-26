package nutsdb

import (
	"github.com/nutsdb/nutsdb/internal/core"
)

// EntryStatus represents the Entry status in the current Tx
type EntryStatus = uint8

const (
	// NotFoundEntry means there is no changes for this entry in current Tx
	NotFoundEntry EntryStatus = 0
	// EntryDeleted means this Entry has been deleted in the current Tx
	EntryDeleted EntryStatus = 1
	// EntryUpdated means this Entry has been updated in the current Tx
	EntryUpdated EntryStatus = 2
)

// BucketStatus represents the current status of bucket in current Tx
type BucketStatus = uint8

const (
	// BucketStatusExistAlready means this bucket already exists
	BucketStatusExistAlready = 1
	// BucketStatusDeleted means this bucket is already deleted
	BucketStatusDeleted = 2
	// BucketStatusNew means this bucket is created in current Tx
	BucketStatusNew = 3
	// BucketStatusUpdated means this bucket is updated in current Tx
	BucketStatusUpdated = 4
	// BucketStatusUnknown means this bucket doesn't exist
	BucketStatusUnknown = 5
)

// pendingBucketList the uncommitted bucket changes in this Tx
type pendingBucketList map[core.Ds]map[core.BucketName]*core.Bucket

// pendingEntriesInBTree means the changes Entries in DataStructureBTree in the Tx
type pendingEntriesInBTree map[core.BucketName]map[string]*core.Entry

// pendingEntryList the uncommitted Entry changes in this Tx
type pendingEntryList struct {
	entriesInBTree pendingEntriesInBTree
	entries        map[core.Ds]map[core.BucketName][]*core.Entry
	size           int
}

// newPendingEntriesList create a new pendingEntryList object for a Tx
func newPendingEntriesList() *pendingEntryList { _ = "STUB: not implemented"; return nil }

// submitEntry submit an entry into pendingEntryList
func (pending *pendingEntryList) submitEntry(ds core.Ds, bucket string, e *core.Entry) {
	_ = "STUB: not implemented"
	return
}

func (pending *pendingEntryList) Get(ds core.Ds, bucket string, key []byte) (entry *core.Entry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pending *pendingEntryList) GetTTL(ds core.Ds, bucket string, key []byte) (ttl int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pending *pendingEntryList) getDataByRange(
	start, end []byte, bucketName core.BucketName,
) (keys, values [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// rangeBucket input a range handler function f and call it with every bucket in pendingBucketList
func (p pendingBucketList) rangeBucket(f func(bucket *core.Bucket) error) error {
	_ = "STUB: not implemented"
	return nil
}

// toList collect all the entries in pendingEntryList to a list.
func (pending *pendingEntryList) toList() []*core.Entry { _ = "STUB: not implemented"; return nil }

func (pending *pendingEntryList) rangeEntries(_ core.Ds, bucketName core.BucketName, rangeFunc func(entry *core.Entry) bool) {
	_ = "STUB: not implemented"
	return
}

func (pending *pendingEntryList) MaxOrMinKey(bucketName string, isMax bool) (key []byte, found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// isBucketNotFoundStatus return true for bucket is not found,
// false for other status.
func isBucketNotFoundStatus(status BucketStatus) bool { _ = "STUB: not implemented"; return false }
