package stream

import "context"

type Collector[T any, U any] interface {
	CollectValue(ctx context.Context, value T, state U) (U, error)
}

type CollectorFunc[T any, U any] func(ctx context.Context, value T, state U) (U, error)

func (fn CollectorFunc[T, U]) CollectValue(ctx context.Context, value T, state U) (U, error) {
	return fn(ctx, value, state)
}

func SimpleCollectorFunc[T any, U any](fn func(value T, state U) U) CollectorFunc[T, U] {
	return CollectorFunc[T, U](func(ctx context.Context, value T, state U) (U, error) {
		return fn(value, state), nil
	})
}
