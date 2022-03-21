package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestSingle(t *testing.T) {
	s := Single(1)
	v, err := s.Next(context.Background())
	require.NoError(t, err)
	require.Equal(t, 1, v)
	v, err = s.Next(context.Background())
	require.Equal(t, io.EOF, err)
	require.Equal(t, 0, v)
}
