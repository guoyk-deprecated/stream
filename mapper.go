package stream

import "context"

type Mapper[T any, U any] func(ctx context.Context, input T) ([]U, error)

func SimpleMapper[T any, U any](fn func(input T) []U) Mapper[T, U] {
	return func(ctx context.Context, input T) ([]U, error) {
		return fn(input), nil
	}
}
