package nutsdb

import (
	"sync"
)

type request struct {
	tx  *Tx
	Wg  sync.WaitGroup
	Err error
	ref int32
}

var requestPool = sync.Pool{
	New: func() any {
		return new(request)
	},
}

func (req *request) reset() { _ = "STUB: not implemented"; return }

func (req *request) IncrRef() { _ = "STUB: not implemented"; return }

func (req *request) DecrRef() { _ = "STUB: not implemented"; return }

func (req *request) Wait() error { _ = "STUB: not implemented"; return nil }

// DecrRef after writing to DB.
