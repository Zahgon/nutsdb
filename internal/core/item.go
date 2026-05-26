package core

type Item[T any] struct {
	Key    []byte
	Record *T
}

func NewItem[T any](key []byte, record *T) *Item[T] { _ = "STUB: not implemented"; return nil }
