package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func (p person) Error() string {
	return fmt.Sprint("Incorrect person")
}

func (p person) hello(name string) error {
	var err error
	if p.Name != name {
		err = p
	} else {
		fmt.Println("Hello ", name)
	}
	return err
}

func main() {
	p1 := person{"James", 87}
	p2 := person{"Bond", 54}

	err := p1.hello("Bond")
	if err != nil {
		fmt.Println(err, err.(person).Name, err.(person).Age)
	}

	people := []person{p1, p2}
	bs, err := json.Marshal(&people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
