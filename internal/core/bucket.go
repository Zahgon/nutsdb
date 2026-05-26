package core

import (
	"errors"

	"github.com/nutsdb/nutsdb/internal/utils"
)

type Ds = uint16
type BucketId = uint64
type BucketName = string
type IDMarkerInBucket map[BucketName]map[Ds]BucketId
type InfoMapperInBucket map[BucketId]*Bucket

var BucketMetaSize int64

const (
	IdSize = 8
	DsSize = 2
)

type BucketOperation uint16

const (
	BucketInsertOperation BucketOperation = 1
	BucketUpdateOperation BucketOperation = 2
	BucketDeleteOperation BucketOperation = 3
)

var ErrBucketCrcInvalid = errors.New("bucket crc invalid")

func init() {
	BucketMetaSize = utils.GetDiskSizeFromSingleObject(BucketMeta{})
}

// BucketMeta stores the Meta info of a Bucket. E.g. the size of bucket it store in disk.
type BucketMeta struct {
	Crc uint32
	// Op: Mark the latest operation (e.g. delete, insert, update) for this bucket.
	Op BucketOperation
	// Size: the size of payload.
	Size uint32
}

// Bucket is the disk structure of bucket
type Bucket struct {
	// Meta: the metadata for this bucket
	Meta *BucketMeta
	// Id: is the marker for this bucket, every bucket creation activity will generate a new Id for it.
	// for example. If you have a bucket called "bucket_1", and you just delete bucket and create it again.
	// the last bucket will have a different Id from the previous one.
	Id BucketId
	// Ds: the data structure for this bucket. (List, Set, SortSet, String)
	Ds Ds
	// Name: the name of this bucket.
	Name string
}

// Decode : CRC | op | size
func (meta *BucketMeta) Decode(bytes []byte) { _ = "STUB: not implemented"; return }

// Encode : Meta | BucketId | Ds | BucketName
func (b *Bucket) Encode() []byte { _ = "STUB: not implemented"; return nil }

// Decode : Meta | BucketId | Ds | BucketName
func (b *Bucket) Decode(bytes []byte) error {
	_ = "STUB: not implemented"
	// parse the payload
	return nil
}

func (b *Bucket) GetEntrySize() int { _ = "STUB: not implemented"; return 0 }

func (b *Bucket) GetCRC(headerBuf []byte, dataBuf []byte) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func (b *Bucket) GetPayloadSize() int { _ = "STUB: not implemented"; return 0 }
