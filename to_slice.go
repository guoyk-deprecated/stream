package stream

func ToSlice[T any]() Collector[T, []T] {
	return SimpleCollector(func(value T, state []T) []T {
		return append(state, value)
	})
}
