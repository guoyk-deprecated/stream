package stream

import "context"

type Mapper[T any, U any] interface {
	MapValue(ctx context.Context, input T) ([]U, error)
}

type MapperFunc[T any, U any] func(ctx context.Context, input T) ([]U, error)

func (fn MapperFunc[T, U]) MapValue(ctx context.Context, input T) ([]U, error) {
	return fn(ctx, input)
}

func SimpleMapperFunc[T any, U any](fn func(input T) []U) MapperFunc[T, U] {
	return func(ctx context.Context, input T) ([]U, error) {
		return fn(input), nil
	}
}
