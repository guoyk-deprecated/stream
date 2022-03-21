package stream

import "context"

type FilterFunc[T any] func(ctx context.Context, value T) (bool, error)

func SimpleFilterFunc[T any](fn func(value T) bool) FilterFunc[T] {
	return func(ctx context.Context, value T) (bool, error) {
		return fn(value), nil
	}
}

func Filter[T any](stream Stream[T], filter FilterFunc[T]) Stream[T] {
	return Flatten(Map(stream, func(ctx context.Context, value T) (Stream[T], error) {
		if ok, err := filter(ctx, value); err != nil {
			return nil, err
		} else {
			if ok {
				return Single(value), nil
			} else {
				return Noop[T](), nil
			}
		}
	}))
}
