package stream

import (
	"context"
	"io"
)

type sliceStream[T any] struct {
	value []T
}

func (s *sliceStream[T]) Next(ctx context.Context) (T, error) {
	var output T
	if len(s.value) == 0 {
		return output, io.EOF
	}
	output = s.value[0]
	s.value = s.value[1:]
	return output, nil
}

func FromSlice[T any](value []T) Stream[T] {
	return &sliceStream[T]{value: value}
}
