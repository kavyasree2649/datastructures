func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	start := 0
	input := []rune(s)

	length := 0
	for i := 0; i < len(input); i++ {

		if index, ok := charMap[input[i]]; ok && index >= start {
			start = index + 1
		}
		charMap[input[i]] = i
		length = max(length, i-start+1)

	}

	return length
}

func max(a, b int) int {
	if a > b {

		return a
	}
	return b
}
