package stream

import (
	"context"
	"io"
	"sync"
)

type flattenStream[T any] struct {
	stream  Stream[Stream[T]]
	lock    sync.Locker
	current Stream[T]
}

func (s *flattenStream[T]) Next(ctx context.Context) (output T, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
redo:
	if s.current == nil {
		if s.current, err = s.stream.Next(ctx); err != nil {
			return
		}
	}
	if output, err = s.current.Next(ctx); err != nil {
		if err == io.EOF {
			s.current = nil
			goto redo
		} else {
			return
		}
	}
	return
}

func Flatten[T any](stream Stream[Stream[T]]) Stream[T] {
	return &flattenStream[T]{stream: stream, lock: &sync.Mutex{}}
}
