// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	input := "This is SO MUCH FUN!"
	slice := strings.Fields(input)
	result := make([]string, 0)

	for _, curr := range slice {
		result = append(result, capitilizeString(curr))
	}

	fmt.Println(strings.Join(result, " "))
}

func capitilizeString(input string) string {
	input = strings.ToLower(input)
	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])
	input = string(runes)
	return input
}
