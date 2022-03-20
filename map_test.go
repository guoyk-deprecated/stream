package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

var countStringConverter = MapperFunc[string, int](func(ctx context.Context, input string) ([]int, error) {
	return []int{len(input)}, nil
})

func TestMap(t *testing.T) {
	upstream := FromSlice([]string{"a", "bb", "ccc"})
	var s Stream[int]
	s = Map[string, int](upstream, countStringConverter)

	ctx := context.Background()

	var v int
	var err error

	v, err = s.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, v)
	v, err = s.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, v)
	v, err = s.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, 3, v)
	v, err = s.Next(ctx)
	require.Error(t, err)
	require.Equal(t, io.EOF, err)
}
