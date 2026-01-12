package main

import "fmt"

// BFS to find shortest distance from src to dest
func shortestDistanceBFS(graph map[int][]int, src, dest int) int {
	queue := []int{src}
	visited := make(map[int]bool)
	distance := make(map[int]int)

	visited[src] = true
	distance[src] = 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Destination reached
		if node == dest {
			return distance[node]
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				distance[neighbor] = distance[node] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	// Destination not reachable
	return -1
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3},
		2: {0, 3},
		3: {1, 2, 4},
		4: {3},
	}

	src := 0
	dest := 4

	fmt.Println("Shortest Distance:", shortestDistanceBFS(graph, src, dest))
}
