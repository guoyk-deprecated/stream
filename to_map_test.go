package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToMap(t *testing.T) {
	s := FromSlice([]Tuple2[int, string]{
		{
			V1: 1,
			V2: "a",
		},
		{
			V1: 2,
			V2: "b",
		},
		{
			V1: 3,
			V2: "c",
		},
	})
	r, err := Collect(
		context.Background(),
		s,
		nil,
		ToMap[int, string](),
	)
	require.NoError(t, err)
	require.Equal(t, map[int]string{1: "a", 2: "b", 3: "c"}, r)
}
