package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 5, 4, 1, 0, 6, 8, 9, 10, 11, 12}

	if len(arr) == 0 {
		fmt.Println(0)
		return
	}

	sort.Ints(arr)

	largest := 1
	cnt := 1

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] == arr[i] {
			continue // skip duplicates
		}
		if arr[i+1] == arr[i]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt > largest {
			largest = cnt
		}
	}

	fmt.Println(largest)
}
