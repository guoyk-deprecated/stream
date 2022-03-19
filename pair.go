package gstream

type Pair[T any, U any] struct {
	V1 T
	V2 U
}

func NewPair[T any, U any](v1 T, v2 U) Pair[T, U] {
	return Pair[T, U]{V1: v1, V2: v2}
}
