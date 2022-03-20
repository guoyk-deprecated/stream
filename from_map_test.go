package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromMap(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	s := FromMap(m)
	r, err := Collect(context.Background(), s, map[int]string{}, ToMap[int, string]())
	require.NoError(t, err)
	require.Equal(t, m, r)
}
