package stream

import "sync"

func ToSlice[T any]() CollectFunc[T, []T] {
	lock := &sync.Mutex{}
	return SimpleCollectFunc(func(value T, state []T) []T {
		lock.Lock()
		defer lock.Unlock()
		return append(state, value)
	})
}
