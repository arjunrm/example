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
	//failedChanUsage()
	correctChanUsage()

	//failedBufChanUsage()
	bufChanUsage()

	// directional channels
	sendOnlyChan()
	//receiveOnlyCha()

	sendReceiveChan()

	usinRangeToReceive()

	usingSelectStatement()

	usingFaninPattern()

	usingFanoutPattern()

	goContextTrace()

	usingGoContext()

	usingGoContext2()
}

func failedChanUsage() {
	c := make(chan int)
	c <- 23
	fmt.Println(<-c)
}

func correctChanUsage() {
	c := make(chan int)
	x := 10

	go func(x *int) {
		*x = 20
		c <- 98
	}(&x)

	fmt.Println(<-c, x)
	fmt.Printf("%T\n", c)
}

func failedBufChanUsage() {
	c := make(chan int)
	c <- 54
	c <- 21
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func bufChanUsage() {
	c := make(chan int, 2)
	c <- 54
	c <- 21
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func sendOnlyChan() {
	c := make(chan int)
	cs := make(chan<- int, 1)
	cs <- 23
	fmt.Printf("%T\n", cs)

	// doesn't work
	// c = cs
	cs = c
}

func receiveOnlyCha() {
	c := make(chan int)
	cr := make(<-chan int)
	fmt.Println(<-cr)

	// doesn't work
	// c = cr
	cr = c
}

func sendReceiveChan() {
	c := make(chan string)

	fmt.Println("sendReceiveChan")

	// send
	go foo(c)

	// receive
	bar(c)
}

func foo(c chan<- string) {
	fmt.Println("--> foo()")
	c <- "Hello Bar"
	fmt.Println("<-- foo()")
}

func bar(c <-chan string) {
	fmt.Println("--> bar()")
	fmt.Println(<-c)
	fmt.Println("<-- bar()")
}

func usinRangeToReceive() {
	c := make(chan int)

	// sender
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	// receiver
	func() {
		// range will continue to pull the values
		// till the channel is closed
		for i := range c {
			fmt.Println(i)
		}
	}()
}

func usingSelectStatement() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go func(eve chan int, odd chan int, quit chan<- int) {
		for i := 0; i < 20; i++ {
			if i%2 == 0 {
				eve <- i
			} else {
				odd <- i
			}
		}
		close(quit)
		//quit <- 99
	}(eve, odd, quit)

	// receive
	func(eve chan int, odd chan int, quit <-chan int) {
		for {
			select {
			case v := <-eve:
				fmt.Println("Even: ", v)
			case v := <-odd:
				fmt.Println("Odd: ", v)
				// when quit chan is closed then ok=false
			case v, ok := <-quit:
				fmt.Println("Quit: ", v, ok)
				return
			}
		}
	}(eve, odd, quit)
}

func usingFaninPattern() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go func(eve chan int, odd chan int) {
		for i := 0; i < 20; i++ {
			if i%2 == 0 {
				eve <- i
			} else {
				odd <- i
			}
		}
		close(eve)
		close(odd)
	}(eve, odd)

	// receive
	go func(eve chan int, odd chan int, fanin chan<- int) {
		var wg sync.WaitGroup
		wg.Add(2)

		// receive even
		go func() {
			for i := range eve {
				fanin <- i
			}
			wg.Done()
		}()

		//receive odd
		go func() {
			for j := range odd {
				fanin <- j
			}
			wg.Done()
		}()

		wg.Wait()
		close(fanin)
	}(eve, odd, fanin)

	// receive fanin
	func(fanin <-chan int) {
		for i := range fanin {
			fmt.Println("Fanin: ", i)
		}
	}(fanin)
}

func usingFanoutPattern() {
	ic := make(chan int)
	oc := make(chan int)

	// populate input channel
	go func(ic chan<- int) {
		for i := 0; i < 10; i++ {
			ic <- i
		}
		close(ic)
	}(ic)

	// fanOut jobs to worker go routines
	go func(ic <-chan int, oc chan<- int) {
		var wg sync.WaitGroup

		for i := range ic {
			wg.Add(1)
			go func(val int, oc chan<- int) {
				// simulate time consuming go routine
				// by sleeping for random amount of time
				oc <- func() int {
					time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
					return val + rand.Intn(500)
				}()
				wg.Done()
			}(i, oc)
		}

		wg.Wait()
		close(oc)
	}(ic, oc)

	// receive from output channel
	for i := range oc {
		fmt.Println("FanOut", i)
	}
}

func goContextTrace() {
	ctx := context.Background()

	fmt.Printf("%v\n", ctx)
	fmt.Printf("%T\n", ctx)
	fmt.Printf("%v\n", ctx.Err())

	doneCtx, cancel := context.WithCancel(ctx)

	fmt.Printf("%v\n", doneCtx)
	fmt.Printf("%T\n", doneCtx)
	fmt.Printf("%v\n", doneCtx.Err())
	fmt.Printf("%v\n", cancel)
	fmt.Printf("%T\n", cancel)

	cancel()

	fmt.Printf("%v\n", doneCtx)
	fmt.Printf("%T\n", doneCtx)
	fmt.Printf("%v\n", doneCtx.Err())
	fmt.Printf("%v\n", cancel)
	fmt.Printf("%T\n", cancel)
}

func usingGoContext() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Ctx error check 1: ", ctx.Err())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done called")
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working on: ", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Ctx error check 2: ", ctx.Err())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	cancel()

	time.Sleep(time.Second * 2)
	fmt.Println("Ctx error check 3: ", ctx.Err())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}

func usingGoContext2() {
	ctx, cancel := context.WithCancel(context.Background())

	// defer func() {
	// 	fmt.Println("Sleeping")
	// 	time.Sleep(time.Second * 1)
	// }()
	defer cancel() // cancel when we are finished

	for v := range gen(ctx) {
		fmt.Println("Range Context : ", v)
		if v == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done called")
				return
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}
