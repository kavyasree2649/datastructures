package main

import (
	"fmt"
)

var cacheMap = make(map[int]int)

func main() {
	out := fibanocii(1)
	fmt.Println("sum: ", out)
}

func fibanocii(n int) int {
	if _, ok := cacheMap[n]; ok {
		return cacheMap[n]
	} else if n < 2 {
		return n
	} else {
		cacheMap[n] = fibanocii(n-1) + fibanocii(n-2)
		return cacheMap[n]
	}

}
