#Convert string to a rune slice

#Traverse only the first half of the string

#Find the first character that is NOT 'a'

#Replace it with 'a'

#Return the result

#If all characters in first half are 'a', change last character to 'b'


package main

import "fmt"

func breakPalindrome(p string) string {
	n := len(p)
	if n == 1 {
		return ""
	}

	runes := []rune(p)

	// Check only first half
	for i := 0; i < n/2; i++ {
		if runes[i] != 'a' {
			runes[i] = 'a'
			return string(runes)
		}
	}

	// All first-half characters are 'a'
	runes[n-1] = 'b'
	return string(runes)
}

func main() {
	fmt.Println(breakPalindrome("abccba")) // aaccba
	fmt.Println(breakPalindrome("a"))      // ""
	fmt.Println(breakPalindrome("aaa"))    // aab
}
