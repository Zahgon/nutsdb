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
	db, _ = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(fileDir),
		nutsdb.WithSegmentSize(1024*1024), // 1MB
	)
	if err != nil {
		panic(err)
	}
	bucket = "bucketForList"
}

func main() {
	testRPushAndLPush()

	testLRange()

	testLPop()

	testRPop()

	testRPushItems()

	testLRem()

	testLRange()

	testLPeek()

	testRPeek()

	testLTrim()

	testLRange()

	testLSize()

	testLRemByIndex()

	testLKeys()
}

func testRPushAndLPush() { _ = "STUB: not implemented"; return }

func testLRange() { _ = "STUB: not implemented"; return }

func testLPop() { _ = "STUB: not implemented"; return }

// val1

func testRPop() { _ = "STUB: not implemented"; return }

// val2

func testRPushItems() { _ = "STUB: not implemented"; return }

func testLRem() { _ = "STUB: not implemented"; return }

// count := 1

func testLPeek() { _ = "STUB: not implemented"; return }

// val11

func testRPeek() { _ = "STUB: not implemented"; return }

// val2

func testLTrim() { _ = "STUB: not implemented"; return }

func testLSize() { _ = "STUB: not implemented"; return }

func testLRemByIndex() { _ = "STUB: not implemented"; return }

func testLKeys() { _ = "STUB: not implemented"; return }

// true: continue, false: break
