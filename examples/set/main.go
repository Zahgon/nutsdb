package main

import (
	"os"

	"github.com/nutsdb/nutsdb"
)

var (
	db     *nutsdb.DB
	bucket string
	err    error
)

func init() {
	fileDir := "/tmp/nutsdb_example"

	files, _ := os.ReadDir(fileDir)
	for _, f := range files {
		name := f.Name()
		if name != "" {
			err := os.RemoveAll(fileDir + "/" + name)
			if err != nil {
				panic(err)
			}
		}
	}
	db, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(fileDir),
		nutsdb.WithSegmentSize(1024*1024), // 1MB
	)
	if err != nil {
		panic(err)
	}
	bucket = "bucketForSet"
}

func main() {
	testSAdd()

	testSAreMembers()

	testSCard()

	testSDiffByOneBucket()

	testSDiffByTwoBuckets()

	testSHasKey()

	testSIsMember()

	testSMembers()

	testSMoveByOneBucket()

	testSMoveByTwoBuckets()

	testSPop()

	testSRem()

	testSUnionByOneBucket()

	testSUnionByTwoBucket()

	testSKeys()
}

func testSAdd() { _ = "STUB: not implemented"; return }

func testSAreMembers() { _ = "STUB: not implemented"; return }

func testSCard() { _ = "STUB: not implemented"; return }

func testSDiffByOneBucket() { _ = "STUB: not implemented"; return }

// item a
// item b

func testSDiffByTwoBuckets() { _ = "STUB: not implemented"; return }

func testSHasKey() { _ = "STUB: not implemented"; return }

func testSIsMember() { _ = "STUB: not implemented"; return }

func testSMembers() { _ = "STUB: not implemented"; return }

func testSMoveByOneBucket() { _ = "STUB: not implemented"; return }

func testSMoveByTwoBuckets() { _ = "STUB: not implemented"; return }

func testSPop() { _ = "STUB: not implemented"; return }

func testSRem() { _ = "STUB: not implemented"; return }

func testSUnionByOneBucket() { _ = "STUB: not implemented"; return }

func testSUnionByTwoBucket() { _ = "STUB: not implemented"; return }

func testSKeys() { _ = "STUB: not implemented"; return }

// true: continue, false: break
