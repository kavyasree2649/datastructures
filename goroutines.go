package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	evenChan := make(chan bool)
	oddChan := make(chan bool)

	n := 10
	wg.Add(2)

	go printOddNumbers(evenChan, oddChan, n, &wg)

	go printEvenNumbers(oddChan, evenChan, n, &wg)

	evenChan <- true

	wg.Wait()
}

func printEvenNumbers(odd, even chan bool, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i += 2 {
		<-even
		fmt.Println("Im even: ", i)

		if i+1 <= n {

			odd <- true
		}
	}
}
func printOddNumbers(even, odd chan bool, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-odd
		fmt.Println("Im odd: ", i)
		if i+1 <= n {
			even <- true

		}
	}
}
