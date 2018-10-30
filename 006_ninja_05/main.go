package main

import (
	"fmt"
)

type person struct {
	first     string
	age       int
	favColors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		first:     "Arya",
		age:       13,
		favColors: []string{"Red", "Green"},
	}
	p2 := person{
		first:     "Jamie",
		age:       40,
		favColors: []string{"Black"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	// map of string:struct
	m1 := map[string]person{}
	m1["Stark"] = p1
	m1["Lannister"] = p2
	fmt.Println(m1)

	for i, v := range m1 {
		fmt.Println(i, v)
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	fmt.Println(t1)

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		luxury: true,
	}
	fmt.Println(s1)

	as1 := struct {
		idMap    map[int]string
		subjects []string
		works    bool
	}{
		idMap: map[int]string{
			1: "Jon",
			2: "Arya",
		},
		subjects: []string{
			"English",
			"Hindi",
		},
		works: true,
	}
	fmt.Println(as1)

	as2 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
		},
		favDrinks: []string{
			"Martini",
		},
	}
	fmt.Println(as2)
}
