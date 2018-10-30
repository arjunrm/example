package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(2)
	go foo()

	go bar()

	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()

	ch := make(chan int)
	go func() {
		ch <- doSomething(4)
	}()
	fmt.Println(<-ch)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}

func doSomething(x int) int {
	return x * 5
}
