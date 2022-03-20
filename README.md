# stream
Golang Streaming API with generics

## Requirements

* Golang 1.18+

## Demo

```go
ctx := context.Background()
s, err := stream.Collect[string, []string](
	ctx,
	stream.Map[int, string](
		stream.FromSlice([]int{1, 2, 3}),
		stream.SimpleMapperFunc(strconv.Itoa),
	),
	nil,
	stream.CollectorToSlice[string](),
)
require.NoError(t, err)
require.Equal(t, []string{"1", "2", "3"}, s)
```

## Credits

Guo Y.K., MIT License