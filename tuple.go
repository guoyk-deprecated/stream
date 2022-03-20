package stream

type Tuple2[T1 any, T2 any] struct {
	V1 T1
	V2 T2
}

func NewTuple2[T1 any, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{V1: v1, V2: v2}
}
