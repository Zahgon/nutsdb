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
	"bytes"
	"sync/atomic"

	"github.com/nutsdb/nutsdb/internal/core"
)

const (
	// txStatusRunning means the tx is running
	txStatusRunning = 1
	// txStatusCommitting means the tx is committing
	txStatusCommitting = 2
	// txStatusClosed means the tx is closed, ether committed or rollback
	txStatusClosed = 3
)

// Tx represents a transaction.
type Tx struct {
	id                uint64
	db                *DB
	writable          bool
	status            atomic.Value
	pendingWrites     *pendingEntryList
	size              int64
	pendingBucketList pendingBucketList
	lockAcquired      atomic.Bool // track if lock was acquired (for WriteBatch optimization)
}

type txnCb struct {
	commit func() error
	user   func(error)
	err    error
}

func (tx *Tx) submitEntry(ds uint16, bucket string, e *core.Entry) {
	_ = "STUB: not implemented"
	return
}

func runTxnCallback(cb *txnCb) { _ = "STUB: not implemented"; return }

// Begin opens a new transaction.
// Multiple read-only transactions can be opened at the same time but there can
// only be one read/write transaction at a time. Attempting to open a read/write
// transactions while another one is in progress will result in blocking until
// the current read/write transaction is completed.
// All transactions must be closed by calling Commit() or Rollback() when done.
func (db *DB) Begin(writable bool) (tx *Tx, err error) {
	_ = "STUB: not implemented"
	// Use TransactionManager so we guard via db.mu when acquiring new transactions
	return nil, nil
}

// newTx returns a newly initialized Tx object at given writable.
func newTx(db *DB, writable bool) (tx *Tx, err error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *Tx) CommitWith(cb func(error)) { _ = "STUB: not implemented"; return }

func (tx *Tx) commitAndSend() (func() error, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *Tx) checkSize() error { _ = "STUB: not implemented"; return nil }

// getTxID returns the tx id.
// Uses cached snowflake node to avoid recreating for every transaction.
func (tx *Tx) getTxID() uint64 { _ = "STUB: not implemented"; return 0 }

// Commit commits the transaction, following these steps:
//
// 1. check the length of pendingWrites.If there are no writes, return immediately.
//
// 2. check if the ActiveFile has not enough space to store entry. if not, call rotateActiveFile function.
//
// 3. write pendingWrites to disk, if a non-nil error,return the error.
//
// 4. build Hint index.
//
// 5. send updated entries to watch manager if watch feature is enabled.
//
// 6. Unlock the database and clear the db field.
func (tx *Tx) Commit() (err error) { _ = "STUB: not implemented"; return nil }

// Ensure the transaction is unregistered so active counts stay accurate

// If the database is closing/closed, abort early to avoid touching released resources.

// judge all write records is whether more than the MaxWriteRecordCount

// add to cache

// send updated entries to watch manager

func (tx *Tx) getNewAddRecordCount() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (tx *Tx) getListEntryNewAddRecordCount(bucketId core.BucketId, entry *core.Entry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) getKvEntryNewAddRecordCount(bucketId core.BucketId, entry *core.Entry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) getSetEntryNewAddRecordCount(_ core.BucketId, entry *core.Entry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) getSortedSetEntryNewAddRecordCount(bucketId core.BucketId, entry *core.Entry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) keyExistsInSortedSet(bucketId core.BucketId, key, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func (tx *Tx) getEntryNewAddRecordCount(entry *core.Entry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) allocCommitBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

// avoid grow

// rotateActiveFile rotates log file when active file is not enough space to store the entry.
func (tx *Tx) rotateActiveFile() error { _ = "STUB: not implemented"; return nil }

// reset ActiveFile

func (tx *Tx) writeData(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Rollback closes the transaction.
func (tx *Tx) Rollback() error { _ = "STUB: not implemented"; return nil }

// Unregister from TransactionManager so active tracking stays correct

// lock locks the database based on the transaction type.
func (tx *Tx) lock() { _ = "STUB: not implemented"; return }

// unlock unlocks the database based on the transaction type.
func (tx *Tx) unlock() { _ = "STUB: not implemented"; return }

// Lock was not acquired, nothing to unlock

func (tx *Tx) handleErr(err error) { _ = "STUB: not implemented"; return }

func (tx *Tx) checkTxIsClosed() error { _ = "STUB: not implemented"; return nil }

// put sets the value for a key in the bucket.
// Returns an error if tx is closed, if performing a write operation on a read-only transaction, if the key is empty.
func (tx *Tx) put(bucket string, key, value []byte, ttl uint32, flag uint16, timestamp uint64, ds uint16) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// setStatusCommitting will change the tx status to txStatusCommitting
func (tx *Tx) setStatusCommitting() { _ = "STUB: not implemented"; return }

// setStatusClosed will change the tx status to txStatusClosed
func (tx *Tx) setStatusClosed() { _ = "STUB: not implemented"; return }

// setStatusRunning will change the tx status to txStatusRunning
func (tx *Tx) setStatusRunning() { _ = "STUB: not implemented"; return }

// isCommitting will check if the tx status is txStatusCommitting
func (tx *Tx) isCommitting() bool { _ = "STUB: not implemented"; return false }

// isClosed will check if the tx status is txStatusClosed
func (tx *Tx) isClosed() bool { _ = "STUB: not implemented"; return false }

func (tx *Tx) buildIdxes(records []*core.Record, entries []*core.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) putBucket(b *core.Bucket) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) SubmitBucket() error { _ = "STUB: not implemented"; return nil }

// buildBucketInIndex build indexes on creation and deletion of buckets
func (tx *Tx) buildBucketInIndex() error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) getChangeCountInEntriesChanges() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tx *Tx) getChangeCountInBucketChanges() int64 { _ = "STUB: not implemented"; return 0 }

// getBucketAndItsStatus, get bucket and it is status in pendingBucketList,
// if bucket is already in bucket manager but not in pendingList, will return BucketStatusExistAlready.
func (tx *Tx) getBucketAndItsStatus(ds core.Ds, name core.BucketName) (BucketStatus, *core.Bucket) {
	_ = "STUB: not implemented"
	return *new(BucketStatus), nil
}

// findEntryStatus finds the latest status for the certain Entry in Tx
func (tx *Tx) findEntryAndItsStatus(_ core.Ds, bucket core.BucketName, key string) (EntryStatus, *core.Entry) {
	_ = "STUB: not implemented"
	return *new(EntryStatus), nil
}

/*
 * send updated entries to watch manager for monitoring
 * and specifying the buckets to be deleted
 * @param pendingWriteList: the list of entries to be sent
 * @param deletedBuckets: the buckets to be deleted
 *
 *
 * @return: nil if success, error if any
 */
func (tx *Tx) sendUpdatedEntries(pendingWriteList []*core.Entry, deletedBuckets map[core.BucketName]bool) {
	_ = "STUB: not implemented"
	return
}

/*
* send buckets to watch manager for specifying the bucket to be deleted

* @param pendingWriteList: the list of entries to be sent
* @return: nil if success, error if any
 */
func (tx *Tx) getDeletedBuckets() (deletedBuckets map[core.BucketName]bool) {
	_ = "STUB: not implemented"
	return nil
}

// registerKeysForActiveExpiration registers keys with TTL in the timing wheel for active expiration.
func (tx *Tx) registerKeysForActiveExpiration(records []*core.Record, entries []*core.Entry) {
	_ = "STUB: not implemented"
	return
}
