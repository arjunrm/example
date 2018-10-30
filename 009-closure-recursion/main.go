package main

import "fmt"

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopfact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	fmt.Println(factorial(5))
	fmt.Println(loopfact(6))
}
