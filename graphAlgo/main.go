package main

import (
	"fmt"
)

func main() {
	graph := make(map[string][]string)

	fmt.Println(graph)

	graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	fmt.Println(graph)

	// depth first travseral
}
