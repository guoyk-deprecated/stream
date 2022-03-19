package stream

import (
	"context"
	"io"
)

type sliceStream[T any] struct {
	idx   int
	value []T
}

func (s *sliceStream[T]) Next(ctx context.Context) (T, error) {
	var val T
	if s.idx >= len(s.value) {
		return val, io.EOF
	}
	val = s.value[s.idx]
	s.idx++
	return val, nil
}

func FromSlice[T any](value []T) Stream[T] {
	return &sliceStream[T]{value: value}
}
