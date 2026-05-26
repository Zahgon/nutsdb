package nutsdb

import (
	"math"
)

const MergeFileIDBase int64 = math.MinInt64

func GetMergeFileID(seq int) int64 { _ = "STUB: not implemented"; return 0 }

func IsMergeFile(fileID int64) bool { _ = "STUB: not implemented"; return false }

func GetMergeSeq(fileID int64) int { _ = "STUB: not implemented"; return 0 }

func mergeFilePrefix() string { _ = "STUB: not implemented"; return "" }

func mergeDataFileName(seq int) string { _ = "STUB: not implemented"; return "" }

func mergeHintFileName(seq int) string { _ = "STUB: not implemented"; return "" }

func getMergeDataPath(dir string, seq int) string { _ = "STUB: not implemented"; return "" }

func getMergeHintPath(dir string, seq int) string { _ = "STUB: not implemented"; return "" }

func parseMergeSeq(name string) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func enumerateDataFileIDs(dir string) (userIDs []int64, mergeIDs []int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func purgeMergeFiles(dir string, keep map[int64]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
