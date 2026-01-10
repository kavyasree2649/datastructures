package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	value int
	index int
}

func minSwapsToSort(arr []int) int {
	n := len(arr)

	// Create value-index pairs
	pairs := make([]Pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = Pair{arr[i], i}
	}

	fmt.Println("1------->", pairs)

	// Sort pairs by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	fmt.Println("2------->", pairs)

	visited := make([]bool, n)
	swaps := 0

	for i := 0; i < n; i++ {
		// Already visited or already in correct position
		if visited[i] || pairs[i].index == i {
			continue
		}

		cycleSize := 0
		j := i

		for !visited[j] {
			visited[j] = true
			j = pairs[j].index
			cycleSize++
		}

		if cycleSize > 1 {
			swaps += cycleSize - 1
		}
	}

	return swaps
}

func main() {
	arr := []int{4, 3, 2, 1}
	fmt.Println(minSwapsToSort(arr)) // Output: 2
}
