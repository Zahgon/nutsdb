package main

import (
	"fmt"
	"os"

	"github.com/nutsdb/nutsdb"
)

var (
	db     *nutsdb.DB
	bucket string
)

func init() {
	fileDir := "/tmp/nutsdb_example"

	files, _ := os.ReadDir(fileDir)
	for _, f := range files {
		name := f.Name()
		if name != "" {
			fmt.Println(fileDir + "/" + name)
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
	bucket = "bucketForString"
}

func main() {
	// create bucket first
	createBucket()

	// insert
	put()
	// read
	read()

	// delete
	delete()
	// read
	read()

	// insert
	put()
	// read
	read()

	// update
	put2()
	// read
	read()

	// get value length
	valueLen()

	// get new value and old value
	getSet()
	// read
	read()

	//put if not exists
	put3()

	//put if exits
	put4()

	// get remaining TTL
	getTTL()

	// save name2 as persistent
	persist()

	// get uncommitted update in same transaction
	getUncommittedUpdateInSameTransaction()
}

func createBucket() { _ = "STUB: not implemented"; return }

func delete() { _ = "STUB: not implemented"; return }

func put() { _ = "STUB: not implemented"; return }

func put2() { _ = "STUB: not implemented"; return }

func put3() { _ = "STUB: not implemented"; return }

func put4() { _ = "STUB: not implemented"; return }

func read() { _ = "STUB: not implemented"; return }

func valueLen() { _ = "STUB: not implemented"; return }

func getSet() { _ = "STUB: not implemented"; return }

func getTTL() { _ = "STUB: not implemented"; return }

func persist() { _ = "STUB: not implemented"; return }

func getUncommittedUpdateInSameTransaction() { _ = "STUB: not implemented"; return }
