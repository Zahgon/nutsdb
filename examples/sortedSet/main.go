package main

import (
	"os"

	"github.com/nutsdb/nutsdb"
)

var (
	db  *nutsdb.DB
	err error
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

}

func main() {
	testZAdd()

	testZScore()

	testZCard()

	testZCount()

	testZMembers()

	testZPeekMax()

	testZPeekMin()

	testZPopMax()

	testZPopMin()

	testZRangeByRank()

	testZRangeByScore()

	testZRank()

	testZRevRank()

	testZRem()

	testZRemRangeByRank()
}

func testZAdd() { _ = "STUB: not implemented"; return }

func testZCard() { _ = "STUB: not implemented"; return }

func testZCount() { _ = "STUB: not implemented"; return }

func testZScore() { _ = "STUB: not implemented"; return }

func testZMembers() { _ = "STUB: not implemented"; return }

func testZPeekMax() { _ = "STUB: not implemented"; return }

// val3

func testZPeekMin() { _ = "STUB: not implemented"; return }

// val1

func testZPopMax() { _ = "STUB: not implemented"; return }

// val3

func testZPopMin() { _ = "STUB: not implemented"; return }

// val1

func testZRangeByRank() { _ = "STUB: not implemented"; return }

func testZRangeByScore() { _ = "STUB: not implemented"; return }

func testZRank() { _ = "STUB: not implemented"; return }

func testZRevRank() { _ = "STUB: not implemented"; return }

// ZRevRank key1 rank: 3

// ZRevRank key2 rank: 2

// ZRevRank key3 rank: 1

func testZRem() { _ = "STUB: not implemented"; return }

func testZRemRangeByRank() { _ = "STUB: not implemented"; return }
