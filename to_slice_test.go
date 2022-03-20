package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCollectorSlice(t *testing.T) {
	upstream := FromSlice([]string{"a", "b", "c"})
	result, err := Collect(context.Background(), upstream, nil, ToSlice[string]())
	require.NoError(t, err)
	require.Equal(t, []string{"a", "b", "c"}, result)
}
