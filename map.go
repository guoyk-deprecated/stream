package stream

import "context"

type Mapper[T any, U any] interface {
	MapValue(ctx context.Context, input T) (U, error)
}

type MapperFunc[T any, U any] func(ctx context.Context, input T) (U, error)

func (fn MapperFunc[T, U]) MapValue(ctx context.Context, input T) (U, error) {
	return fn(ctx, input)
}

func SimpleMapperFunc[T any, U any](fn func(input T) U) MapperFunc[T, U] {
	return func(ctx context.Context, input T) (U, error) {
		return fn(input), nil
	}
}

type mapStream[T any, U any] struct {
	upstream  Stream[T]
	converter Mapper[T, U]
}

func (s *mapStream[T, U]) Next(ctx context.Context) (output U, err error) {
	var input T
	if input, err = s.upstream.Next(ctx); err != nil {
		return
	}
	if output, err = s.converter.MapValue(ctx, input); err != nil {
		return
	}
	return
}

func Map[T any, U any](upstream Stream[T], converter Mapper[T, U]) Stream[U] {
	return &mapStream[T, U]{
		upstream:  upstream,
		converter: converter,
	}
}
