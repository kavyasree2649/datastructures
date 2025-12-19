package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4}
	n := len(arr)
	largest := 0

	for i := 1; i < n-2; {

		//check if peak or not
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {

			cnt := 1
			j := i

			//count backwards
			for j >= 1 && arr[j] > arr[j-1] {
				j--
				cnt++

			}

			//count forwards

			for i <= n-2 && arr[i] > arr[i+1] {
				i++
				cnt++
			}

			if cnt > largest {
				largest = cnt
			}

		} else {
			//increment i if not peak
			i++
		}

	}

	fmt.Println(largest)

}
