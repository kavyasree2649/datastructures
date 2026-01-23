// Frequency count using binary search
package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 1, 2, 2, 2, 2, 2, 3, 4, 4, 5}
	key := 2

	frequencyCount := upper_bound(arr, key) - lower_bound(arr, key) + 1
	fmt.Println(frequencyCount)
}

func lower_bound(arr []int, key int) int {

	s := 0

	e := len(arr) - 1

	ans := -1

	for s <= e {

		mid := (s + e) / 2

		if arr[mid] == key {
			ans = mid
			e = mid - 1
		} else if arr[mid] > key {

			e = mid - 1

		} else {
			s = mid + 1
		}

	}

	return ans
}

func upper_bound(arr []int, key int) int {
	s := 0

	e := len(arr) - 1

	ans := -1

	for s <= e {

		mid := (s + e) / 2

		if arr[mid] == key {
			ans = mid
			s = mid + 1
		} else if arr[mid] > key {

			e = mid - 1

		} else {
			s = mid + 1
		}

	}

	return ans

}
