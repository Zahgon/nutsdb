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

package core

const Persistent uint32 = 0

// Record means item of indexes in memory
type Record struct {
	Key       []byte
	Value     []byte
	FileID    int64
	DataPos   uint64
	ValueSize uint32
	Timestamp uint64
	TTL       uint32
	TxID      uint64
}

// NewRecord generate a record Obj
func NewRecord() *Record { _ = "STUB: not implemented"; return nil }

func (r *Record) WithKey(k []byte) *Record { _ = "STUB: not implemented"; return nil }

// WithValue set the Value to Record
func (r *Record) WithValue(v []byte) *Record { _ = "STUB: not implemented"; return nil }

// WithFileId set FileID to Record
func (r *Record) WithFileId(fid int64) *Record { _ = "STUB: not implemented"; return nil }

// WithDataPos set DataPos to Record
func (r *Record) WithDataPos(pos uint64) *Record { _ = "STUB: not implemented"; return nil }

func (r *Record) WithValueSize(valueSize uint32) *Record { _ = "STUB: not implemented"; return nil }

func (r *Record) WithTimestamp(timestamp uint64) *Record { _ = "STUB: not implemented"; return nil }

func (r *Record) WithTTL(ttl uint32) *Record { _ = "STUB: not implemented"; return nil }

func (r *Record) WithTxID(txID uint64) *Record { _ = "STUB: not implemented"; return nil }
