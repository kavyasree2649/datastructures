package main

import (
	"fmt"
	"strings"
)

//Splits by any whitespace (space, tab, newline)

func main() {
	text := "hello   world  golang"
	tokens := strings.Fields(text)

	fmt.Println(tokens)
}




package main

import (
	"fmt"
	"strings"
)

//Split by a specific delimiter

func main() {
	text := "apple,banana,orange"
	tokens := strings.Split(text, ",")

	fmt.Println(tokens)
}
