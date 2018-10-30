package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for []user
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	us1 := user{"James", 45}
	us2 := user{"MoneyPenny", 35}

	client := []user{us1, us2}
	bs, err := json.Marshal(client)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	var serviceProviders []user
	err = json.Unmarshal(bs, &serviceProviders)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(serviceProviders)
	}

	for _, sp := range serviceProviders {
		fmt.Println(sp.First, sp.Age)
	}

	json.NewEncoder(os.Stdout).Encode(client)
	fmt.Println(json.Valid(bs))

	xi := []int{124, 543, 524, 4546348, 879780, 31, 25, 98}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	xs := []string{"asft", "fbfihno", "acvreh", "zmner", "pkohr"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	sort.Sort(ByAge(client))
	fmt.Println(client)
}
