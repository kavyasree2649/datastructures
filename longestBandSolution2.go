package main

import (
	"fmt"
)

// program to find longest band
func main() {

	arr := []int{2, 5, 4, 1, 0, 6, 8, 9, 10, 11, 12}

	largest := 1

	lookUpMap := make(map[int]bool)

	for _, val := range arr {
		if _, ok := lookUpMap[val]; !ok {
			lookUpMap[val] = true
		}
	}
	for _, val := range arr {
		if !lookUpMap[val-1] {

			length := 1
			next := val + 1

			for lookUpMap[next] {
				length++
				next++
			}

			if length > largest {
				largest = length
			}
		}

	}

	fmt.Println(largest)
}
