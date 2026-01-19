// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func MergeSort(arr1 []int, arr2 []int) []int {

	i := 0
	j := 0
	var finalResp []int
	for i < len(arr1)-1 && j < len(arr2)-1 {

		if arr1[i] < arr2[j] {
			finalResp = append(finalResp, arr1[i])
			i++
		} else {
			finalResp = append(finalResp, arr2[j])
			j++
		}

	}

	fmt.Println(arr1[i+1:], arr2[j:])

	finalResp = append(finalResp, arr1[i:]...)
	finalResp = append(finalResp, arr2[j:]...)
	return finalResp
}

func main() {
	arr1 := []int{2, 4, 6}
	arr2 := []int{3, 9, 10}

	fmt.Println(MergeSort(arr1, arr2))
}
