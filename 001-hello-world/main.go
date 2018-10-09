package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var i int
	i = 100
	str := "Arjun Ramamurthy "
	fmt.Printf("%d\n", i)
	fmt.Println(str)

	for i := 0; i < 3; i++ {
		sampleFunc(i)
	}
}

func sampleFunc(i int) {
	fmt.Println("Sample func", i)
}

func sampleFunc2(i float64) {
	fmt.Println("Sample func 2")
}
