package gstream

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestSlice(t *testing.T) {
	var v string
	var err error
	ctx := context.Background()
	s := FromSlice([]string{"a", "b", "c"})
	v, err = s.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, "a", v)
	v, err = s.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, "b", v)
	v, err = s.Next(ctx)
	require.NoError(t, err)
	require.Equal(t, "c", v)
	v, err = s.Next(ctx)
	require.Error(t, err)
	require.Equal(t, io.EOF, err)

	var input []string
	s = FromSlice(input)
	v, err = s.Next(ctx)
	require.Error(t, err)
	require.Equal(t, io.EOF, err)
}
