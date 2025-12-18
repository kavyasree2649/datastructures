package main

import (
	"fmt"
)

func main() {
	input := []int{5, 6, 9, 4}

	targetSum := 10

	complementMap := make(map[int]bool, 0)
	for _, val := range input {

		complement := targetSum - val

		if _, ok := complementMap[complement]; ok {
			fmt.Println(val, complement)
			return
		}

		complementMap[val] = true

	}

	fmt.Println(complementMap)

}
