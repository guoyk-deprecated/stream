package stream

import (
	"context"
	"io"
	"sync"
)

type sliceStream[T any] struct {
	value []T
	lock  sync.Locker
}

func (s *sliceStream[T]) Next(ctx context.Context) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var output T
	if len(s.value) == 0 {
		return output, io.EOF
	}
	output = s.value[0]
	s.value = s.value[1:]
	return output, nil
}

func FromSlice[T any](value []T) Stream[T] {
	if len(value) == 0 {
		return Noop[T]()
	}
	if len(value) == 1 {
		return Single(value[0])
	}
	return &sliceStream[T]{value: value, lock: &sync.Mutex{}}
}

func Literal[T any](value ...T) Stream[T] {
	return FromSlice(value)
}
