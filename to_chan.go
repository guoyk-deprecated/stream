package stream

func ToChan[T any]() CollectFunc[T, chan T] {
	return SimpleCollectFunc(func(value T, state chan T) chan T {
		state <- value
		return state
	})
}
