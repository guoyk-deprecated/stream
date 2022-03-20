package stream

func FromMap[T comparable, U any](value map[T]U) Stream[Tuple2[T, U]] {
	var input []Tuple2[T, U]
	for k, v := range value {
		input = append(input, NewTuple2(k, v))
	}
	return FromSlice(input)
}
