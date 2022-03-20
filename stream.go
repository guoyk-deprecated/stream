package stream

import "context"

type Stream[T any] interface {
	Next(ctx context.Context) (T, error)
}

type Func[T any] func(ctx context.Context) (T, error)

func (f Func[T]) Next(ctx context.Context) (T, error) {
	return f(ctx)
}
