package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToChan(t *testing.T) {
	s := Literal(1, 2, 3)
	out := make(chan int, 10)
	r, err := Collect(context.Background(), s, out, ToChan[int]())
	require.NoError(t, err)
	require.Equal(t, 1, <-r)
	require.Equal(t, 2, <-r)
	require.Equal(t, 3, <-r)
}
