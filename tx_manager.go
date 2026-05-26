package nutsdb

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nutsdb/nutsdb/internal/core"
)

type txManager struct {
	lifecycle core.ComponentLifecycle

	db            *DB
	statusManager *StatusManager

	activeTxs     sync.Map
	activeTxCount atomic.Int64

	maxActiveTxs int64
	configMu     sync.RWMutex
}

func newTxManager(db *DB, sm *StatusManager) *txManager { _ = "STUB: not implemented"; return nil }

func (tm *txManager) Name() string { _ = "STUB: not implemented"; return "" }

func (tm *txManager) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (tm *txManager) Stop(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

// acquireLock is false when WriteBatch already holds db.mu.
func (tm *txManager) BeginTx(writable bool, acquireLock bool) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tm *txManager) RegisterTx(tx *Tx) error { _ = "STUB: not implemented"; return nil }

func (tm *txManager) UnregisterTx(txID uint64) { _ = "STUB: not implemented"; return }

func (tm *txManager) GetActiveTxCount() int64 { _ = "STUB: not implemented"; return 0 }

func (tm *txManager) WaitForActiveTxs(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// 0 means unlimited.
func (tm *txManager) SetMaxActiveTxs(max int64) { _ = "STUB: not implemented"; return }

func (tm *txManager) GetMaxActiveTxs() int64 { _ = "STUB: not implemented"; return 0 }

func (tm *txManager) GetActiveTxs() []uint64 { _ = "STUB: not implemented"; return nil }

// forceAbortActiveTxs unregisters and closes all active transactions.
// Used during shutdown when transactions fail to finish within the timeout.
func (tm *txManager) forceAbortActiveTxs() { _ = "STUB: not implemented"; return }
