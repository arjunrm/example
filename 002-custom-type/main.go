package main

import (
	"fmt"
	"runtime"
)

type dog int

var a int
var b dog
var flag bool

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 56
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// conversions
	a = int(b)
	c := 25.0
	fmt.Println(c)
	d := int(c)
	fmt.Println(d)

	flag = true
	fmt.Printf("%v\t%T\n", flag, flag)
	var x rune // alias for int32
	fmt.Printf("%v\t%T\n", x, x)
	var y byte // alias for uint8
	fmt.Printf("%v\t%T\n", y, y)

	fmt.Println(runtime.GOARCH, runtime.GOOS, runtime.Compiler)
}
