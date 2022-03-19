package stream

import "context"

func CollectorToSlice[T any]() Collector[T, []T] {
	return CollectorFunc[T, []T](func(ctx context.Context, value T, state []T) ([]T, error) {
		return append(state, value), nil
	})
}
