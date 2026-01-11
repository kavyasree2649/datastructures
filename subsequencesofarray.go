package main

import "fmt"

func subsequences(arr []int, index int, curr []int) {
	if index == len(arr) {
		fmt.Println(curr)
		return
	}

	// 1️⃣ Include current element
	subsequences(arr, index+1, append(curr, arr[index]))

	// 2️⃣ Exclude current element
	subsequences(arr, index+1, curr)
}

func main() {
	arr := []int{1, 2, 3}
	subsequences(arr, 0, []int{})
}
