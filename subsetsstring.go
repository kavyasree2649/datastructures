package main

import "fmt"

func main() {
	input := "hello world"
	subset := "yllo"

	s := []rune(input)
	t := []rune(subset)

	i := len(s) - 1
	j := len(t) - 1

	for i >= 0 && j >= 0 {
		if s[i] == t[j] {
			i--
			j--
		} else {
			i--
		}
	}

	fmt.Println(j == -1)
}
