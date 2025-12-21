package main

import (
	"fmt"
)

// program to rains collected
func main() {

	h := []int{0, 1, 2, 1, 2, 3, 0, 1, 1, 2}

	n := len(h)

	if n <= 2 {
		fmt.Println(0)
		return
	}

	//make left and right arrary
	left := make([]int, len(h))
	right := make([]int, len(h))

	//populate left array

	left[0] = h[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], h[i])
	}

	//populate right array
	right[n-1] = h[n-1]

	for i := 1; i < n; i++ {
		right[n-i-1] = max(right[n-i], h[n-i-1])
	}
	water := 0
	for i := 0; i < n; i++ {
		water += min(left[i], right[i]) - h[i]
	}

	fmt.Println(water)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
