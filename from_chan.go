package stream

import (
	"context"
	"io"
)

type chanStream[T any] struct {
	value <-chan T
}

func (s chanStream[T]) Next(ctx context.Context) (T, error) {
	var (
		v  T
		ok bool
	)
	if v, ok = <-s.value; !ok {
		return v, io.EOF
	} else {
		return v, nil
	}
}

func FromChan[T any](ch <-chan T) Stream[T] {
	return chanStream[T]{value: ch}
}
