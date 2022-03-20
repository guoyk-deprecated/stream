package stream

import (
	"context"
	"io"
)

func Collect[T any, U any](ctx context.Context, upstream Stream[T], first U, collector Collector[T, U]) (output U, err error) {
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
