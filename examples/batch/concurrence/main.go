package main

import (
	"github.com/nutsdb/nutsdb"
)

func readWorker(id int, jobs <-chan int, results chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func writeWorker(id int, jobs <-chan int, results chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

var (
	db     *nutsdb.DB
	err    error
	bucket string
)

func removeFileDir(fileDir string) { _ = "STUB: not implemented"; return }

func main() {
	removeFlag := false
	fileDir := "/tmp/nutsdb_example_concurrence"

	if removeFlag {
		removeFileDir(fileDir)
	}

	db, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(fileDir),
		nutsdb.WithSegmentSize(1024*1024), // 1MB
	)
	if err != nil {
		panic(err)
	}

	bucket = "bucketForString"
	// create bucket first
	createBucket()

	readJobs := make(chan int, 10)
	writeJobs := make(chan int, 10)
	readResults := make(chan struct{}, 10)
	writeResults := make(chan struct{}, 10)

	for w := 1; w <= 3; w++ {
		go readWorker(w, readJobs, readResults)
	}

	for w := 1; w <= 3; w++ {
		go writeWorker(w, writeJobs, writeResults)
	}

	for j := 1; j <= 10; j++ {
		readJobs <- j
	}

	for j := 1; j <= 10; j++ {
		writeJobs <- j
	}

	close(readJobs)
	close(writeJobs)

	for a := 1; a <= 10; a++ {
		<-readResults
	}
	for a := 1; a <= 10; a++ {
		<-writeResults
	}
}

func createBucket() { _ = "STUB: not implemented"; return }
