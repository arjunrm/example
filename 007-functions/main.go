package main

import "fmt"

// function/method syntax
// func (r receiver) identifier(params) (return(s)) { code }

// (r receiver) - used to attach the func to a type
// any value of that type will have access to this func

type person struct {
	first string
	last  string
}

func (speaker person) speak() string {
	return fmt.Sprint("Method of struct person")
}

func main() {
	defer foo()
	bar()

	// everything in go is pass by value
	s1 := retString("James")
	fmt.Println(s1)

	age, name := mulReturns()
	fmt.Println(age, name)

	fmt.Println(sum("sum1", 1, 2, 3, 4, 5, 6))

	si := []int{4, 3, 5, 6, 7, 3, 1, 7}
	// unfurling a slice
	// si... - in this case new slice is not created
	// same slice it passed by value i.e. same underlying array
	fmt.Println(sum("sum2", si...))
	// variadic params accepts 0 or more values
	// it should be the final param
	fmt.Println(sum(""))
	fmt.Println("Jon")

	p1 := person{
		first: "Jamie",
		last:  "Lannister",
	}
	fmt.Println(p1, p1.speak())
}

func retString(str string) string {
	return fmt.Sprint("Hello ", str)
}

func mulReturns() (int, string) {
	return 10, "Jon"
}

// variadic params are slice of type Ex: []int
func sum(name string, x ...int) (string, int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return name, sum
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
