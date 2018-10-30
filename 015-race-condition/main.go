package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//raceCondition()
	//noRaceConditionMutex()
	noRaceConditionAtomic()
}

func raceCondition() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// yields the processor for other go routines to run
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Goroutines: ", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func noRaceConditionMutex() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mut sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mut.Lock()
			v := counter
			// yields the processor for other go routines to run
			runtime.Gosched()
			v++
			counter = v
			mut.Unlock()
			fmt.Println("Goroutines: ", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func noRaceConditionAtomic() {
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
