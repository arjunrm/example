package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// SortByAge is used for sorting person by age
type SortByAge []person

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	fmt.Println()

	p1 := person{First: "James", Last: "Bond", Age: 25}
	p2 := person{"Miss", "Moneypenny", 27}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	var people2 []person
	err2 := json.Unmarshal(bs, &people2)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println(people2)

	str := `[{"First":"Jon","Last":"Snow","Age":35},{"First":"James","Last":"Bond","Age":27}]`
	byteSlice := []byte(str)

	var people3 []person
	err3 := json.Unmarshal(byteSlice, &people3)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Printf("%+v\n", people3)

	// io.Writer interface
	// anyone implementing func Write(data []byte) (n int, err error)
	// is of type Writer
	fmt.Println()
	io.WriteString(os.Stdout, "Hello")

	fmt.Println()

	// sorting
	si1 := []int{235, 6, 457, 235, 567, 236, 1}
	fmt.Println(si1)
	sort.Ints(si1)
	fmt.Println(si1)

	ss1 := []string{"James", "asf", "asfgrbeh", "Dr. No"}
	fmt.Println(ss1)
	sort.Strings(ss1)
	fmt.Println(ss1)

	// custom sort
	fmt.Println(people3)
	sort.Sort(SortByAge(people3))
	fmt.Println(people3)

	pwd := "password1234"
	fmt.Println(pwd)
	pwdbs, errpwd := bcrypt.GenerateFromPassword([]byte(pwd), 5)
	if errpwd != nil {
		fmt.Println(errpwd)
	}
	fmt.Println(pwdbs)

	retErr := bcrypt.CompareHashAndPassword(pwdbs, []byte(pwd))
	if retErr != nil {
		fmt.Println(retErr)
	} else {
		fmt.Println("Logged in!!")
	}
}
