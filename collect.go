package stream

import (
	"context"
	"io"
)

type CollectFunc[T any, U any] func(ctx context.Context, value T, state U) (U, error)

func SimpleCollectFunc[T any, U any](fn func(value T, state U) U) CollectFunc[T, U] {
	return func(ctx context.Context, value T, state U) (U, error) {
		return fn(value, state), nil
	}
}

func Collect[T any, U any](ctx context.Context, upstream Stream[T], first U, collector CollectFunc[T, U]) (output U, err error) {
	var value T
	output = first
	for {
		if value, err = upstream.Next(ctx); err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return
			}
		}
		if output, err = collector(ctx, value, output); err != nil {
			return
		}
	}
	return
}
