package main

import "fmt"

//Control flow:
// 1. sequence
// 2. loop/iterative
// 3. conditional

var x = 10

func main() {
	fmt.Println("Hello World")
	var i int
	i = 100
	// short declaration operator
	j := 100.00
	str := "Arjun Ramamurthy "
	fmt.Printf("%d %f, %T %T\n", i, j, i, j)
	fmt.Println(str)

	y := 25
	fmt.Printf("%v\t%b\t%x\t%#x\n", y, y, y, y)

	for i := 0; i < 3; i++ {
		sampleFunc(i)
		sampleFunc2(float64(i))
		sampleFunc3("str")
	}

	z := fmt.Sprintf("sample value: %v", x)
	fmt.Println(z)
}

func sampleFunc(i int) {
	fmt.Println("Sample func 1:", i)
}

func sampleFunc2(i float64) {
	fmt.Println("Sample func 2:", i)
}

func sampleFunc3(str string) {
	n, err := fmt.Println("Sample func 3:", str)
	fmt.Println(n, err)
}
