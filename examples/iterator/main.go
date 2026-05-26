package main

import (
	"fmt"

	"github.com/nutsdb/nutsdb"
)

var (
	db     *nutsdb.DB
	bucket = "bucket_iterator_demo"
)

func init() {
	db, _ = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("/tmp/nutsdbexample/example_iterator"),
	)
}

func main() {
	tx, err := db.Begin(true)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		key := []byte("key_" + fmt.Sprintf("%03d", i))
		val := []byte("val_" + fmt.Sprintf("%03d", i))
		if err = tx.Put(bucket, key, val, nutsdb.Persistent); err != nil {
			// tx rollback
			_ = tx.Rollback()
			fmt.Printf("rollback ok, err %v:", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	// forward iteration
	forwardIteration()
	// reverse iterative
	reverseIterative()
}

func forwardIteration() { _ = "STUB: not implemented"; return }

func reverseIterative() { _ = "STUB: not implemented"; return }
