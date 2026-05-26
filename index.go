// Copyright 2022 The nutsdb Author. All rights reserved.
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
	"github.com/nutsdb/nutsdb/internal/core"
	"github.com/nutsdb/nutsdb/internal/data"
)

type IndexType interface {
	data.BTree | data.Set | SortedSet | data.List
}

type defaultOp[T IndexType] struct {
	Idx map[core.BucketId]*T
}

func (op *defaultOp[T]) computeIfAbsent(id core.BucketId, f func() *T) *T {
	_ = "STUB: not implemented"
	return nil
}

func (op *defaultOp[T]) delete(id core.BucketId) { _ = "STUB: not implemented"; return }

func (op *defaultOp[T]) exist(id core.BucketId) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (op *defaultOp[T]) getIdxLen() int { _ = "STUB: not implemented"; return 0 }

func (op *defaultOp[T]) rangeIdx(f func(elem *T)) { _ = "STUB: not implemented"; return }

type Index struct {
	List      *ListIndex
	BTree     *BTreeIndex
	Set       *SetIndex
	SortedSet *SortedSetIndex
	db        *DB
}

func (db *DB) newIndex() *Index { _ = "STUB: not implemented"; return nil }

type ListIndex struct {
	*defaultOp[data.List]
	index *Index
}

func (idx *ListIndex) GetWithDefault(id core.BucketId) *data.List {
	_ = "STUB: not implemented"
	return nil
}

type BTreeIndex struct {
	*defaultOp[data.BTree]
	index *Index
}

func (idx *BTreeIndex) GetWithDefault(id core.BucketId) *data.BTree {
	_ = "STUB: not implemented"
	return nil
}

type SetIndex struct {
	*defaultOp[data.Set]
	index *Index
}

func (idx *SetIndex) GetWithDefault(id core.BucketId) *data.Set {
	_ = "STUB: not implemented"
	return nil
}

type SortedSetIndex struct {
	*defaultOp[SortedSet]
	index *Index
}

func (idx *SortedSetIndex) GetWithDefault(id core.BucketId) *SortedSet {
	_ = "STUB: not implemented"
	return nil
}
