package stream

import (
	"context"
	"io"
)

type Stream[T any] interface {
	Next(ctx context.Context) (T, error)
}

type Func[T any] func(ctx context.Context) (T, error)

func (f Func[T]) Next(ctx context.Context) (T, error) {
	return f(ctx)
}

type noopStream[T any] struct{}

func (s noopStream[T]) Next(ctx context.Context) (output T, err error) {
	err = io.EOF
	return
}

func Noop[T any]() Stream[T] {
	return noopStream[T]{}
}
