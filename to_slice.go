package stream

func ToSlice[T any]() Collector[T, []T] {
	return SimpleCollectorFunc[T, []T](func(value T, state []T) []T {
		return append(state, value)
	})
}
