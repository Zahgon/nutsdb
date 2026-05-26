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
	"time"
)

// getDataPath returns the data path for the given file ID.
func getDataPath(fID int64, dir string) string { _ = "STUB: not implemented"; return "" }

func splitIntStringStr(str, separator string) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func splitStringIntStr(str, separator string) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

func splitIntIntStr(str, separator string) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func decodeListKey(buf []byte) ([]byte, uint64) { _ = "STUB: not implemented"; return nil, 0 }

func splitStringFloat64Str(str, separator string) (string, float64) {
	_ = "STUB: not implemented"
	return "", 0
}

func createNewBufferWithSize(size int) *bytes.Buffer { _ = "STUB: not implemented"; return nil }

// compareAndReturn use bytes.Compare(other, target), if return value is
// comVal, return other, else return target.
func compareAndReturn(target []byte, other []byte, cmpVal int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func expireTime(timestamp uint64, ttl uint32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func mergeKeyValues(
	k0, v0 [][]byte,
	k1, v1 [][]byte,
) (keys, values [][]byte) {
	_ = "STUB: not implemented"
	return nil,

		// Pre-allocate capacity with estimated maximum size to reduce slice re-growth
		nil
}

// skip k1 item if k0 == k1

// This Type is for sort a pair of (k, v).
type sortkv struct {
	k, v [][]byte
}

func (skv *sortkv) Len() int { _ = "STUB: not implemented"; return 0 }

func (skv *sortkv) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (skv *sortkv) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
