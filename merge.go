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
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/nutsdb/nutsdb/internal/core"
)

var ErrDontNeedMerge = errors.New("the number of files waiting to be merged is less than 2")

func (db *DB) Merge() error { _ = "STUB: not implemented"; return nil }

func (db *DB) merge() error { _ = "STUB: not implemented"; return nil }

// Merge removes dirty data and reduce data redundancy,following these steps:
//
// 1. Filter delete or expired entry.
//
// 2. Write entry to activeFile if the key not exist，if exist miss this write operation.
//
// 3. Filter the entry which is committed.
//
// 4. At last remove the merged files.
//
// Caveat: merge is Called means starting multiple write transactions, and it
// will affect the other write request. so execute it at the appropriate time.
func (db *DB) mergeLegacy() error { _ = "STUB: not implemented"; return nil }

// Index reads can race with commits; use a transaction to avoid inconsistencies.

func (db *DB) buildHintFilesAfterMerge(startFileID, endFileID int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) isPendingMergeEntry(entry *core.Entry) bool { _ = "STUB: not implemented"; return false }

func (db *DB) isPendingBtreeEntry(entry *core.Entry) bool { _ = "STUB: not implemented"; return false }

func (db *DB) isPendingSetEntry(entry *core.Entry) bool { _ = "STUB: not implemented"; return false }

func (db *DB) isPendingZSetEntry(entry *core.Entry) bool { _ = "STUB: not implemented"; return false }

func (db *DB) isPendingListEntry(entry *core.Entry) bool { _ = "STUB: not implemented"; return false }

type mergedEntryInfo struct {
	entry    *core.Entry
	bucketId core.BucketId
}

type mergeWorker struct {
	lifecycle core.ComponentLifecycle

	db            *DB
	statusManager *StatusManager

	mergeStartCh chan struct{}
	mergeEndCh   chan error
	isMerging    atomic.Bool

	ticker *time.Ticker

	config MergeConfig
}

type MergeConfig struct {
	MergeInterval   time.Duration
	EnableAutoMerge bool
}

func DefaultMergeConfig() MergeConfig { _ = "STUB: not implemented"; return *new(MergeConfig) }

func newMergeWorker(db *DB, sm *StatusManager, config MergeConfig) *mergeWorker {
	_ = "STUB: not implemented"
	return nil
}

func (mw *mergeWorker) Name() string { _ = "STUB: not implemented"; return "" }

func (mw *mergeWorker) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (mw *mergeWorker) Stop(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

func (mw *mergeWorker) TriggerMerge() error { _ = "STUB: not implemented"; return nil }

func (mw *mergeWorker) IsMerging() bool { _ = "STUB: not implemented"; return false }

func (mw *mergeWorker) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (mw *mergeWorker) performMerge() error { _ = "STUB: not implemented"; return nil }

func (mw *mergeWorker) SetMergeInterval(interval time.Duration) { _ = "STUB: not implemented"; return }

func (mw *mergeWorker) GetMergeInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
