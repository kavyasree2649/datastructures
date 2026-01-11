// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	input := "abcd"

	subsequences(input, "", 0)

}

func subsequences(input string, current string, index int) {

	if index == len(input) {
		fmt.Println(current)
		return
	}

	// include
	subsequences(input, current+string(input[index]), index+1)

	// exlucde
	subsequences(input, current, index+1)

}
