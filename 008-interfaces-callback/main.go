package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

// attching method for struct person
func (p person) speak() string {
	return fmt.Sprintf("%v %v is %v yrs old", p.first, p.last, p.age)
}

func (sa secretAgent) speak() string {
	return fmt.Sprintf("%v %v is %v yrs old and ltk(%v)", sa.first, sa.last, sa.age, sa.ltk)
}

// anybody having method speak is also of type human
type human interface {
	speak() string
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Type person:", h.(person).first)
	case secretAgent:
		fmt.Println("Type secretAgent:", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func emptyInterface(i ...interface{}) {
	fmt.Println(i...)
}

func retFunc() func(int) string {
	return func(x int) string {
		return fmt.Sprint("Returning a func: ", x)
	}
}

func sum(si ...int) int {
	sum := 0
	for _, v := range si {
		sum += v
	}
	return sum
}

func sumCallback(sum int, cb func(bool)) {
	if sum < 10 {
		cb(true)
	} else {
		cb(false)
	}
}

func sumlt10(flag bool) {
	fmt.Println("Sumlt10 cb: ", flag)
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func main() {
	p1 := person{
		first: "Dr.",
		last:  "No",
		age:   35,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.speak())

	fmt.Println(sa1)
	fmt.Println(sa1.speak())

	bar(p1)
	bar(sa1)

	// conversion
	var h1 human
	h1 = p1
	fmt.Println("\nUsing var of type interface")
	fmt.Println(h1.speak())

	i := 10
	j := 31.5
	fmt.Println("\nEmpty Interface")
	emptyInterface(p1.speak())
	emptyInterface(sa1.speak())
	emptyInterface(i, j)

	func() {
		fmt.Println("Anonymous function")
	}()

	func(x int) {
		fmt.Println("Anonymous function: ", x)
	}(10)

	// func expression
	f1 := func() {
		fmt.Println("Func expression")
	}
	f1()

	f2 := retFunc()
	fmt.Println(f2(25))

	fmt.Println(retFunc()(14))

	sumCallback(sum(1, 2, 3), sumlt10)
	sumCallback(sum(5, 6, 7, 8), sumlt10)
	ii := []int{1, 2, 3, 4, 5, 6, 7}
	sumCallback(sum(ii...), sumlt10)

	fmt.Println()
	fmt.Println(even(sum, ii...))
}
