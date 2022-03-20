package stream

import "context"

type Collector[T any, U any] func(ctx context.Context, value T, state U) (U, error)

func SimpleCollector[T any, U any](fn func(value T, state U) U) Collector[T, U] {
	return func(ctx context.Context, value T, state U) (U, error) {
		return fn(value, state), nil
	}
}
