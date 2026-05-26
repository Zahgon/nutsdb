package utils

import (
	"hash/fnv"
)

var fnvHash = fnv.New32a()

func GetFnv32(value []byte) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ConvertBigEndianBytesToUint64(data []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func ConvertUint64ToBigEndianBytes(value uint64) []byte { _ = "STUB: not implemented"; return nil }

func MarshalInts(ints []int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func UnmarshalInts(data []byte) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func MatchForRange(pattern, bucket string, f func(bucket string) bool) (end bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func UvarintSize(x uint64) int { _ = "STUB: not implemented"; return 0 }

func VarintSize(x int64) int { _ = "STUB: not implemented"; return 0 }

func GetDiskSizeFromSingleObject(obj any) int64 { _ = "STUB: not implemented"; return 0 }

// Currently, we only use the unsigned value type for our metadata.go. That's reasonable for us.
// Because it's not possible to use negative value mark the size of data.
// But if you want to make it more flexible, please help yourself.

func OneOfUint16Array(value uint16, array []uint16) bool { _ = "STUB: not implemented"; return false }

func EncodeListKey(key []byte, seq uint64) []byte { _ = "STUB: not implemented"; return nil }

func DecodeListKey(buf []byte) ([]byte, uint64) { _ = "STUB: not implemented"; return nil, 0 }
