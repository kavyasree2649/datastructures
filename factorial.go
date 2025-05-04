package main

import (
	"fmt"
)

func main() {

	resp := factorial(5)
	fmt.Println(resp)

}

func factorial(num int) int {

	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
