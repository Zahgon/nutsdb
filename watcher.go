package nutsdb

import (
	"sync"
	"time"
)

type WatchingFunc func() error

type Watcher struct {
	readyCh      chan struct{} // the channel to signal that the watcher is ready
	watchingFunc WatchingFunc  // the function to watch the key and bucket
	isReady      bool          // indicates whether the watcher is ready

	muReady sync.Mutex
}

func (w *Watcher) WaitReady(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

func (w *Watcher) Run() error { _ = "STUB: not implemented"; return nil }
