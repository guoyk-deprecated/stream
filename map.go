package stream

import "context"

type mapStream[T any, U any] struct {
	stream Stream[T]
	mapper Mapper[T, U]
	cache  []U
}

func (s *mapStream[T, U]) Next(ctx context.Context) (output U, err error) {
	if len(s.cache) != 0 {
		output = s.cache[0]
		s.cache = s.cache[1:]
		return
	}
	var input T
retry:
	if input, err = s.stream.Next(ctx); err != nil {
		return
	}
	if s.cache, err = s.mapper(ctx, input); err != nil {
		return
	}
	if len(s.cache) == 0 {
		goto retry
	}
	output = s.cache[0]
	s.cache = s.cache[1:]
	return
}

func Map[T any, U any](stream Stream[T], mapper Mapper[T, U]) Stream[U] {
	return &mapStream[T, U]{
		stream: stream,
		mapper: mapper,
	}
}
