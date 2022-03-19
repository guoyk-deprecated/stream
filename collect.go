package stream

import (
	"context"
	"io"
)

type Collector[T any, U any] interface {
	CollectValue(ctx context.Context, value T, state U) (U, error)
}

type CollectorFunc[T any, U any] func(ctx context.Context, value T, state U) (U, error)

func (fn CollectorFunc[T, U]) CollectValue(ctx context.Context, value T, state U) (U, error) {
	return fn(ctx, value, state)
}

func Collect[T any, U any](ctx context.Context, upstream Stream[T], collector Collector[T, U]) (result U, err error) {
	var value T
	for {
		if value, err = upstream.Next(ctx); err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return
			}
		}
		if result, err = collector.CollectValue(ctx, value, result); err != nil {
			return
		}
	}
	return
}
