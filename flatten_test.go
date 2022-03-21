package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlatten(t *testing.T) {
	s1 := Literal(1, 2, 3)
	s2 := Literal(4, 5, 6)
	s3 := Literal(7, 8, 9)
	s := Flatten(Literal(s1, s2, s3))
	r, err := Collect(context.Background(), s, nil, ToSlice[int]())
	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, r)
}
