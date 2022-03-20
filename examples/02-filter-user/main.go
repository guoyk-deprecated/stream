package main

import (
	"context"
	"github.com/guoyk93/stream"
	"log"
	"strings"
)

type User struct {
	Name   string
	Active bool
}

func main() {
	// mock a database
	var database = map[string][]User{
		"section-a": {
			{Name: "alice", Active: false}, {Name: "bob", Active: true},
		},
		"section-b": {
			{Name: "alex", Active: true},
		},
	}

	// build Stream[string] of sections
	sections := stream.Literal("section-a", "section-b")
	// get Stream[User] by query 'database'
	users := stream.Map(
		sections,
		func(ctx context.Context, section string) ([]User, error) {
			// this is just a mock, you can do actual database query here
			return database[section], nil
		},
	)
	// get Stream[string] of names by filter and map User.Name
	names := stream.Map(
		users,
		// SimpleMapper is just a wrapper to ignore ctx and error
		stream.SimpleMapper(func(u User) []string {
			if u.Active && strings.HasPrefix(u.Name, "a") {
				return []string{u.Name}
			} else {
				return nil
			}
		}),
	)
	// collect Stream[string] as []string
	// remember, this is when actual operations are executed
	result, _ := stream.Collect(
		context.Background(),
		names,
		nil,
		stream.ToSlice[string](),
	)
	log.Println(result)
	// ["alex"]
}
