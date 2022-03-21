package stream

func FromMap[T comparable, U any](value map[T]U) Stream[Tuple2[T, U]] {
	if len(value) == 0 {
		return Noop[Tuple2[T, U]]()
	}
	input := make([]Tuple2[T, U], 0, len(value))
	for k, v := range value {
		input = append(input, NewTuple2(k, v))
	}
	return FromSlice(input)
}
