package gstream

func FromMap[T comparable, U any](value map[T]U) Stream[Pair[T, U]] {
	var input []Pair[T, U]
	for k, v := range value {
		input = append(input, NewPair(k, v))
	}
	return FromSlice(input)
}
