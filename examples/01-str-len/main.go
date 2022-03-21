package main

import (
	"context"
	"github.com/guoyk93/stream"
	"log"
)

func main() {
	s := stream.Map(
		stream.Literal("a", "bb", "ccc"),
		stream.SimpleMapFunc(func(input string) []int {
			return []int{len(input)}
		}),
	)
	r, _ := stream.Collect(context.Background(), s, nil, stream.ToSlice[int]())
	log.Println(r)
}
