package nutsdb

import (
	"errors"
	"os"

	"github.com/nutsdb/nutsdb/internal/core"
)

var ErrBucketNotExist = errors.New("bucket not found")

const BucketStoreFileName = "bucket.Meta"

type BucketManager struct {
	fd *os.File
	// BucketInfoMapper BucketID => Bucket itself
	BucketInfoMapper core.InfoMapperInBucket
	BucketIDMarker   core.IDMarkerInBucket
	// IDGenerator helps generates an ID for every single bucket
	Gen *IDGenerator
}

func NewBucketManager(dir string) (*BucketManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bucketSubmitRequest struct {
	ds     core.Ds
	name   core.BucketName
	bucket *core.Bucket
}

func (bm *BucketManager) SubmitPendingBucketChange(reqs []*bucketSubmitRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// update the marker info

// recover maxid otherwise new bucket start from 1 again

type IDGenerator struct {
	currentMaxId uint64
}

func (g *IDGenerator) GenId() uint64 { _ = "STUB: not implemented"; return 0 }

func (g *IDGenerator) CompareAndSetMaxId(id uint64) { _ = "STUB: not implemented"; return }

func (bm *BucketManager) ExistBucket(ds core.Ds, name core.BucketName) bool {
	_ = "STUB: not implemented"
	return false
}

func (bm *BucketManager) GetBucket(ds core.Ds, name core.BucketName) (b *core.Bucket, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bm *BucketManager) GetBucketById(id core.BucketId) (*core.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bm *BucketManager) GetBucketID(ds core.Ds, name core.BucketName) (core.BucketId, error) {
	_ = "STUB: not implemented"
	return *new(core.BucketId), nil
}

func (bm *BucketManager) Close() error { _ = "STUB: not implemented"; return nil }
