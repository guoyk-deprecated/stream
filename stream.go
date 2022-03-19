package stream

import "context"

type Stream[T any] interface {
	Next(ctx context.Context) (T, error)
}

type StreamFunc[T any] func(ctx context.Context) (T, error)

func (s StreamFunc[T]) Next(ctx context.Context) (T, error) {
	return s(ctx)
}
