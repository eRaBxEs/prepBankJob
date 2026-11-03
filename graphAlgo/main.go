package main

import (
	"fmt"
)

func main() {
	// create some map structure inorder to use an adjacency list to represent the graph
	graph, graph2 := make(map[string][]string), make(map[string][]string)

	fmt.Println(graph)
	fmt.Println(graph2)

	graph = map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	graph2 = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	fmt.Println("Depth First Traversal (iterative):")
	depthFirstTraversal(graph, "a")
	fmt.Println("===================================")

	fmt.Println("Depth First Traversal (recursive):")
	// depthFirstRecursive(graph2, "a")
	fmt.Println()

	fmt.Println("Breadth First Traversal:")
	// breadthFirstTraversal(graph, "a")

}

// depth first traversal: iterative DFS using a stack
func depthFirstTraversal(graph map[string][]string, source string) {
	stack := []string{source}

	for len(stack) > 0 {
		// pop from the end of the stack: removes the last item of that array
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(current)

		for _, neighbour := range graph[current] {
			stack = append(stack, neighbour)
		}
	}

}

func depthFirstRecursive(graph map[string][]string, source string) {
	fmt.Println(source)
	for _, neighbour := range graph[source] {
		depthFirstRecursive(graph, neighbour)
	}
}
