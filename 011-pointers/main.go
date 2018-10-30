package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	fmt.Println("Before", p.name, p.age)

	p.name += "_changed"
	p.age += 10

	fmt.Println("After", p.name, p.age)
}

func main() {
	a := 42
	fmt.Println(a, &a)
	fmt.Printf("%T\t%T\n", a, &a)

	b := &a
	fmt.Println(b)
	fmt.Println(b, &b, *b)

	p1 := person{"James", 35}
	fmt.Println(p1)

	changeMe(&p1)
}
