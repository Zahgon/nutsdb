package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nutsdb/nutsdb"
)

var (
	db     *nutsdb.DB
	bucket = "bucket_watcher_demo"
	dir    = "/tmp/nutsdbexample/example_watcher"
)

func init() {
	var err error
	os.RemoveAll(dir)

	db, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(dir),
		nutsdb.WithEnableWatch(true),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create bucket first
	if err := db.Update(func(tx *nutsdb.Tx) error {
		return tx.NewBucket(nutsdb.DataStructureBTree, bucket)
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== NutsDB Watcher Example ===")

	// basic watching
	basicWatchExample()

	// watch with multiple operations
	multipleOperationsExample()

	// watch with TTL expiration
	watchWithTTLExample()

	if err := db.Close(); err != nil {
		log.Printf("Error closing database: %v\n", err)
	}
}

func basicWatchExample() { _ = "STUB: not implemented"; return }

// signal that we received the message

// wait for the watcher to be ready

func multipleOperationsExample() { _ = "STUB: not implemented"; return }

// Perform multiple operations

//manually cancel the watcher

func watchWithTTLExample() { _ = "STUB: not implemented"; return }

// we must call the view to trigger the ttl expiration
