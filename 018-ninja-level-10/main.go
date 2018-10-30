package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	rangeOverChan()

	receiveUsingSelectAndTimeout()

	calculateFibonacci()
}

func rangeOverChan() {
	c := gen()

	receive(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receiveUsingSelectAndTimeout() {
	eve := make(chan int)
	odd := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)

	// send
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 2)
			v := rand.Intn(500)
			if v%2 == 0 {
				eve <- v
			} else {
				odd <- v
			}
		}
		close(eve)
		close(odd)
		cancel()
	}()

	// receive
	func() {
		for {
			select {
			case v := <-eve:
				fmt.Println("Even: ", v)
			case v := <-odd:
				fmt.Println("Odd: ", v)
			case <-ctx.Done():
				fmt.Println("Done")
				return
			}
		}
	}()
}

type fiboContainer struct {
	index  int
	series []int
}

func calculateFibonacci() {
	ic := make(chan int)
	oc := make(chan fiboContainer)
	var wg sync.WaitGroup
	wg.Add(3)

	// put values into input chan
	go input(ic, &wg)

	// generate fibonaccci series
	go genFibo(ic, oc, &wg)

	go dispFibo(oc, &wg)

	fmt.Println("calculateFibonacci(): ", runtime.NumGoroutine())

	wg.Wait()
}

func input(ic chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		ic <- i
	}
	close(ic)
	wg.Done()
}

func genFibo(ic <-chan int, oc chan<- fiboContainer, wg *sync.WaitGroup) {
	var fibowg sync.WaitGroup

	// fibo generator
	for v := range ic {
		fibowg.Add(1)
		// create go routine for handling each fibonacci series generation
		go func(v int) {
			var fibo []int
			for i := 1; i <= v; i++ {
				// slow it down
				time.Sleep(time.Millisecond * 500)
				fibo = append(fibo, fibonaccci(i))
			}
			oc <- fiboContainer{v, fibo}
			fibowg.Done()
		}(v)
	}

	fmt.Println("genFibo(): ", runtime.NumGoroutine())

	fibowg.Wait()
	close(oc)
	wg.Done()
}

func dispFibo(oc <-chan fiboContainer, wg *sync.WaitGroup) {
	for v := range oc {
		fmt.Printf("Index: %v, Fibo: %v\n", v.index, v.series)
	}
	wg.Done()
}

func fibonaccci(x int) int {
	if x <= 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fibonaccci(x-1) + fibonaccci(x-2)
}
