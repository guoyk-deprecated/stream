package stream

import "context"

type mapStream[T any, U any] struct {
	upstream  Stream[T]
	converter Mapper[T, U]
}

func (s *mapStream[T, U]) Next(ctx context.Context) (output U, err error) {
	var input T
	if input, err = s.upstream.Next(ctx); err != nil {
		return
	}
	if output, err = s.converter.MapValue(ctx, input); err != nil {
		return
	}
	return
}

func Map[T any, U any](upstream Stream[T], converter Mapper[T, U]) Stream[U] {
	return &mapStream[T, U]{
		upstream:  upstream,
		converter: converter,
	}
}
