package stream

import (
	"context"
	"io"
	"sync/atomic"
)

type singleStram[T any] struct {
	value T
	done  uint32
}

func (s *singleStram[T]) Next(ctx context.Context) (T, error) {
	if atomic.CompareAndSwapUint32(&s.done, 0, 1) {
		output := s.value
		return output, nil
	} else {
		var value T
		return value, io.EOF
	}
}

func Single[T any](value T) Stream[T] {
	return &singleStram[T]{value: value}
}
