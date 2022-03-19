# gstream
Golang Streaming API with generics

## Requirements

* Golang 1.18+

## Demo

```go
ctx := context.Background()
s, err := gstream.Collect[string, []string](
	ctx,
	gstream.Map[int, string](
		gstream.FromSlice([]int{1, 2, 3}),
		gstream.SimpleMapperFunc(strconv.Itoa),
	),
	gstream.CollectorToSlice[string](),
)
require.NoError(t, err)
require.Equal(t, []string{"1", "2", "3"}, s)
```

## Credits

Guo Y.K., MIT License