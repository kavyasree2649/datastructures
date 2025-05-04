package main

import (
	"fmt"
)

func main() {
	input := []int{2, 5, 1, 99, 45, 89, 100, 56}
	fmt.Println(selectionSort(input))
}

func selectionSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}

	return nums
}
