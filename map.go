package stream

import (
	"context"
	"sync"
)

type MapFunc[T any, U any] func(ctx context.Context, input T) ([]U, error)

func SimpleMapFunc[T any, U any](fn func(input T) []U) MapFunc[T, U] {
	return func(ctx context.Context, input T) ([]U, error) {
		return fn(input), nil
	}
}

type mapStream[T any, U any] struct {
	stream Stream[T]
	mapper MapFunc[T, U]
	lock   sync.Locker
	cache  []U
}

func (s *mapStream[T, U]) Next(ctx context.Context) (output U, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.cache) != 0 {
		output = s.cache[0]
		s.cache = s.cache[1:]
		return
	}
	var input T
redo:
	if input, err = s.stream.Next(ctx); err != nil {
		return
	}
	if s.cache, err = s.mapper(ctx, input); err != nil {
		return
	}
	if len(s.cache) == 0 {
		goto redo
	}
	output = s.cache[0]
	s.cache = s.cache[1:]
	return
}

func Map[T any, U any](stream Stream[T], mapper MapFunc[T, U]) Stream[U] {
	return &mapStream[T, U]{
		stream: stream,
		mapper: mapper,
		lock:   &sync.Mutex{},
	}
}
