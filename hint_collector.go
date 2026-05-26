package nutsdb

import "errors"

var errHintCollectorClosed = errors.New("hint collector closed")

const DefaultHintCollectorFlushEvery = 1024

type hintWriter interface {
	Write(*HintEntry) error
	Sync() error
	Close() error
}

type HintCollector struct {
	writer     hintWriter
	buf        []HintEntry
	fileID     int64
	flushEvery int
	closed     bool
}

func NewHintCollector(fileID int64, writer hintWriter, flushEvery int) *HintCollector {
	_ = "STUB: not implemented"
	return nil
}

func (hc *HintCollector) Add(entry *HintEntry) error { _ = "STUB: not implemented"; return nil }

func (hc *HintCollector) Flush() error { _ = "STUB: not implemented"; return nil }

func (hc *HintCollector) Sync() error { _ = "STUB: not implemented"; return nil }

func (hc *HintCollector) Close() error { _ = "STUB: not implemented"; return nil }

func (hc *HintCollector) flush(sync bool) error { _ = "STUB: not implemented"; return nil }
