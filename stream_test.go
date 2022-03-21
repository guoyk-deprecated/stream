package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestNoop(t *testing.T) {
	s := Noop[string]()
	v, err := s.Next(context.Background())
	require.Equal(t, io.EOF, err)
	require.Equal(t, "", v)
}
