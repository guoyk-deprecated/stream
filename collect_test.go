package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCollect(t *testing.T) {
	s := Literal(1, 2, 3)
	c := 0
	_, err := Collect(context.Background(), s, struct{}{}, func(ctx context.Context, value int, state struct{}) (struct{}, error) {
		c += value
		return state, nil
	})
	require.NoError(t, err)
	require.Equal(t, 6, c)
}
