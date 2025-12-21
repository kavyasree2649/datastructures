package main

import "fmt"

func main() {
	str := "abc"
	n := len(str)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Println(str[i:j])
		}
	}
}
