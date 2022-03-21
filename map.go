package stream

import (
	"context"
)

type MapFunc[T any, U any] func(ctx context.Context, input T) (U, error)

func SimpleMapFunc[T any, U any](fn func(input T) U) MapFunc[T, U] {
	return func(ctx context.Context, input T) (U, error) {
		return fn(input), nil
	}
}

type mapStream[T any, U any] struct {
	stream Stream[T]
	mapper MapFunc[T, U]
}

func (s *mapStream[T, U]) Next(ctx context.Context) (output U, err error) {
	var input T
	if input, err = s.stream.Next(ctx); err != nil {
		return
	}
	if output, err = s.mapper(ctx, input); err != nil {
		return
	}
	return
}

func Map[T any, U any](stream Stream[T], mapper MapFunc[T, U]) Stream[U] {
	return &mapStream[T, U]{
		stream: stream,
		mapper: mapper,
	}
}
