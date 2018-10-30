package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type person struct {
	name string
	age  int
}

type perError struct {
	err string
}

func (p *perError) Error() string {
	return p.err
}

func (p person) Hello(name string) error {
	pe := perError{}
	if p.name != name {
		pe.err = "Incorrect name: " + name
	} else {
		fmt.Println("I am James Bond")
	}
	return &pe
}

var logf *os.File

func main() {
	setLoggingToFile()
	defer logf.Close()

	errorHandlingSample()

	//createFile()

	readFile()

	panicExample()

	infoErrors()
}

func errorHandlingSample() {
	p1 := person{"James", 35}

	err := p1.Hello("Bond")
	if err != nil {
		fmt.Printf("%T\n", err)
		fmt.Println(err)
	}

	err = p1.Hello("James")
	if err != nil {
		fmt.Println(err)
	}
}

func createFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")
	n, err := io.Copy(f, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Written: ", n)
}

func setLoggingToFile() {
	logf, err := os.OpenFile("log.txt", os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logf)
}

func readFile() {
	f, err := os.Open("name.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f: ", r)
		}
	}()
	fmt.Println("Caling g")
	g(0)
	fmt.Println("Returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g: ", i)
	fmt.Println("Printing in g: ", i)
	g(i + 1)
}

func infoErrors() {
	v, err := func() (int, error) {
		return 23, errors.New("Sample error info")
	}()
	fmt.Println(v, err)
	fmt.Printf("%T\n", err)

	v, err = func() (int, error) {
		i := 10
		return 54, fmt.Errorf("fmt formatted error: %v", i)
	}()
	fmt.Println(v, err)
	fmt.Printf("%T\n", err)
}
