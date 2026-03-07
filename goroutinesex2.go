package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	signalChan1 := make(chan int)
	signalChan2 := make(chan int)
	signalChan3 := make(chan int)

	n := 10

	wg.Add(3)

	go routine1(n, signalChan1, signalChan2, &wg)
	go routine2(n, signalChan2, signalChan3, &wg)
	go routine3(n, signalChan3, signalChan1, &wg)

	signalChan1 <- 1

	wg.Wait()
}

func routine1(n int, sig1, sig2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-sig1
		if !ok {
			close(sig2)
			return
		}

		if val > n {
			close(sig2)
			return
		}

		fmt.Println("goroutine1:", val)
		sig2 <- val + 1
	}
}

func routine2(n int, sig2, sig3 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-sig2
		if !ok {
			close(sig3)
			return
		}

		if val > n {
			close(sig3)
			return
		}

		fmt.Println("goroutine2:", val)
		sig3 <- val + 1
	}
}

func routine3(n int, sig3, sig1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-sig3
		if !ok {
			close(sig1)
			return
		}

		if val > n {
			// sig1 <- val
			return
		}

		fmt.Println("goroutine3:", val)
		sig1 <- val + 1
	}
}
