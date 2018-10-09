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
		sampleFunc2(float64(i))
		sampleFunc3("str")
	}
}

func sampleFunc(i int) {
	fmt.Println("Sample func 1:", i)
}

func sampleFunc2(i float64) {
	fmt.Println("Sample func 2:", i)
}

func sampleFunc3(str string) {
	fmt.Println("Sample func 3:", str)
}
