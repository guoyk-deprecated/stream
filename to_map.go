package stream

func ToMap[T comparable, U any]() Collector[Tuple2[T, U], map[T]U] {
	return SimpleCollector(func(value Tuple2[T, U], state map[T]U) map[T]U {
		state[value.V1] = value.V2
		return state
	})
}
