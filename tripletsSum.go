package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{1, 2, 4, 5, 6, 9, 4, 10, 3, 7, 8}
	targetSum := 10

	var finalInput [][]int

	sort.Ints(input)
	n := len(input)

	for i := 0; i < n-2; i++ {

		j := i + 1
		k := n - 1
		for j < k {
			currentSum := input[i]
			currentSum += input[j]
			currentSum += input[k]

			if currentSum == targetSum {
				finalInput = append(finalInput, []int{input[i], input[j], input[k]})
				j++
				k--

			} else if currentSum > targetSum {
				k--
			} else if currentSum < targetSum {
				j++
			}
		}

	}

	fmt.Println(finalInput)

}
