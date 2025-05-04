package main

import (
	"fmt"
)

func main() {
	input := []int{2, 5, 1, 99, 45, 89, 100, 56}
	fmt.Println(bubbleSort(input))
}

func bubbleSort(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
