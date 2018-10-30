package main

import (
	"fmt"
	"math"
)

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func endOfMain() {
	fmt.Println("End of main")
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func addFunc() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func callback(xi []int) int {
	if len(xi) == 0 {
		return 0
	} else if len(xi) == 1 {
		return 1
	} else {
		return xi[0] + xi[len(xi)-1]
	}
}

func fooCallback(f func(xi []int) int, ii []int) int {
	n := f(ii)
	fmt.Println("fooCallback")
	fmt.Println(n)
	n++
	return n
}

func main() {
	defer endOfMain()
	defer func() {
		fmt.Println("Defered anonymous func")
	}()

	ii := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	fmt.Println(foo(ii...))

	cir := circle{
		radius: 12.345,
	}
	sq := square{
		length: 14,
	}
	info(cir)
	info(sq)

	fn1 := addFunc()
	fmt.Println(fn1(5, 10))

	si := []int{1, 2, 3, 4, 5}
	fmt.Println(fooCallback(callback, si))
}
