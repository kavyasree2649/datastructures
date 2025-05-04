package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(fibbanocci(i))
	}

}

// 0,1,1,2,3,5
func fibbanocci(num int) int {

	if num < 2 {
		return num
	}
	return fibbanocci(num-2) + fibbanocci(num-1)
}
