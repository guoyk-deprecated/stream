package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter(t *testing.T) {
	s := Filter(Literal(1, 2, 3, 4), SimpleFilterFunc(func(value int) bool {
		return value%2 == 0
	}))
	v, err := Collect(context.Background(), s, nil, ToSlice[int]())
	require.NoError(t, err)
	require.Equal(t, []int{2, 4}, v)
}
