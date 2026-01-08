package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(a []int) []int {
	n := len(a)

	// Make a copy of the array
	b := make([]int, n)
	copy(b, a)

	// Sort the copied array
	sort.Ints(b)

	i := 0
	for i < n && a[i] == b[i] {
		i++
	}

	// If already sorted
	if i == n {
		return []int{-1, -1}
	}

	j := n - 1
	for j >= 0 && a[j] == b[j] {
		j--
	}

	return []int{i, j}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 7, 6, 8, 9, 10}
	result := findUnsortedSubarray(a)
	fmt.Println(result)
}
