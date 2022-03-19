package stream

import (
	"context"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestStream(t *testing.T) {
	ctx := context.Background()
	s, err := Collect[string, []string](
		ctx,
		Map[int, string](
			FromSlice([]int{1, 2, 3}),
			SimpleMapperFunc(strconv.Itoa),
		),
		CollectorToSlice[string](),
	)
	require.NoError(t, err)
	require.Equal(t, []string{"1", "2", "3"}, s)
}
