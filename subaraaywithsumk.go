func subarraysWithSumK(arr []int, k int) {
	i, sum := 0, 0

	for j := 0; j < len(arr); j++ {
		sum += arr[j]

		for sum > k {
			sum -= arr[i]
			i++
		}

		if sum == k {
			fmt.Println(arr[i : j+1])
		}
	}
}
