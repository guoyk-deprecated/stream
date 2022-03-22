package stream

import (
	"context"
	"io"
	"sync/atomic"
	"unsafe"
)

type singleStram[T any] struct {
	value unsafe.Pointer
}

func (s *singleStram[T]) Next(ctx context.Context) (value T, err error) {
	if v := atomic.SwapPointer(&s.value, nil); v != nil {
		value = *(*T)(v)
		return
	} else {
		err = io.EOF
		return
	}
}

func Single[T any](value T) Stream[T] {
	return &singleStram[T]{value: unsafe.Pointer(&value)}
}
