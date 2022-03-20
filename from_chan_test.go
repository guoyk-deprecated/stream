package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFromChan(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	s := FromChan[int](ch)
	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()
	r, err := Collect(
		context.Background(),
		s,
		nil,
		ToSlice[int](),
	)
	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3}, r)
}
