func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	i := 0
	j := 0
	var finalArray []int

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			finalArray = append(finalArray, nums2[j])
			j++
		} else {

			finalArray = append(finalArray, nums1[i])
			i++
		}

	}

	finalArray = append(finalArray, nums1[i:]...)
	finalArray = append(finalArray, nums2[j:]...)

	if len(finalArray)%2 == 0 {

		mid := len(finalArray) / 2

		return float64(finalArray[mid] + finalArray[mid-1] )/ 2.0

	} else {

		median := len(finalArray) / 2

		return float64(finalArray[median])
	}

}
