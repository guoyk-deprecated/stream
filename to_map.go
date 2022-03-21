package stream

import "sync"

func ToMap[T comparable, U any]() CollectFunc[Tuple2[T, U], map[T]U] {
	lock := &sync.Mutex{}
	return SimpleCollectFunc(func(value Tuple2[T, U], state map[T]U) map[T]U {
		lock.Lock()
		defer lock.Unlock()
		state[value.V1] = value.V2
		return state
	})
}
